package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
	"github.com/nats-io/stan.go"
)

const (
	host      = "localhost"
	port      = "5432"
	user      = "postgres"
	password  = "1"
	dbname    = "test"
	tableName = "data"
	natsTheme = "test"
	natsURL   = stan.DefaultNatsURL
	clusterID = "test-cluster"
	clientID  = "0"
	idValue   = "order_uid"
)

var dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
var upgrader = websocket.Upgrader{
	//уменьшим размер буфера до 2048 байт и будем проверять источник запоса по url, принимая запорсы только с данного адреса
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://127.0.0.1:8080"
	},
}

type Data struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func (data *Data) set(key string, val []byte) bool {
	var ok bool
	data.mu.Lock()
	defer data.mu.Unlock()
	if _, ok = data.data[key]; !ok {
		data.data[key] = val
	} else {
		fmt.Printf("we already have data with id %v\n", key)
	}
	return !ok
}
func (data *Data) get(key string) []byte {
	data.mu.RLock() //при чтении будем использовать rlock, чтобы не блокровать map для чтения из других горунтин
	defer data.mu.RUnlock()
	val, ok := data.data[key]
	if !ok {
		val = ([]byte)(fmt.Sprintf("element with key %v not found", key)) //
	}
	return val
}

func main() {

	var wg sync.WaitGroup
	data := &Data{data: make(map[string][]byte)}

	err := getAllDataDB(data) //все полученые данные будут писаться в data
	if err != nil {
		return
	}
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		fmt.Printf("nats-streaming server connect error: %v\n", err)
		return
	}
	fmt.Printf(`connect to nats-streaming server: 
	URL: %v
	channel: %v
	clusterID: %v 
	clientID: %v`+"\n", natsURL, natsTheme, clusterID, clientID)

	ch := make(chan stan.Msg, 100)
	///
	defer sc.Close()
	defer wg.Wait()
	defer close(ch)
	///
	_, err = sc.Subscribe(natsTheme,
		func(m *stan.Msg) {
			ch <- *m
		},
		stan.AckWait(5*time.Second))

	if err != nil {
		fmt.Printf("nats subscribe error: %v\v", err)
		return
	}
	// для ускорения обработки данных из ch запустим обрабботку поступащих данных в горунтинах
	for i := 0; i < runtime.NumCPU()-1; i++ {
		wg.Add(1)
		go processMsg(ch, data, &wg)
	}
	//обработчик запорсов для шаблона /request - запрос с таким шаблоном отправляется при нажатии кнопки "send request"
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("upgrader func error %v\n", err)
			return
		}
		defer conn.Close()
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("read message error %v\n", err)
				return
			}
			id := string(msg)
			res := data.get(id)
			fmt.Printf("%s sent: %v\n", conn.RemoteAddr(), id)
			if err = conn.WriteMessage(msgType, res); err != nil {
				fmt.Printf("write message error %v\n", err)
				return
			}
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html") //websockets.html содержит шаблон страницы
	})
	http.ListenAndServe("localhost:8080", nil)
}

func processMsg(in <-chan stan.Msg, data *Data, wg *sync.WaitGroup) error {

	defer wg.Done()
	defer fmt.Println("end go")

	for {
		var result map[string]interface{}
		msg, ok := <-in
		byteValue := msg.Data
		if !ok {
			break
		}
		err := json.Unmarshal([]byte(byteValue), &result)
		if err != nil {
			return err
		}
		key, ok := result[idValue].(string) //из json достанем id, сами данные отсавим хранить в []bytes
		if ok {
			ok := data.set(key, byteValue)
			if ok {
				err := setDataDB(dbURL, key, byteValue)
				if err != nil {
					return err
				}
			}
		} else {
			fmt.Printf("msg %d from channel %v dont have id\n", msg.Sequence, msg.Subject)
		}
	}
	return nil
}
func getAllDataDB(data *Data) error {

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		return err
	}
	fmt.Printf("Connect to DB: %s\n", dbname)
	defer fmt.Println("close DB connection")
	defer conn.Close(context.Background())

	var rows pgx.Rows
	query := "SELECT * FROM " + tableName + ";"
	rows, err = conn.Query(context.Background(), query)

	if err != nil {
		fmt.Printf("bad get query error: \n%v", err)
		return err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			fmt.Printf("bad get query rows: \n%v", err)
			return err
		}
		id := (string)(values[0].(string))
		val := ([]byte)(values[1].(string))
		data.set(id, val)
	}
	return nil
}
func setDataDB(url, key string, val []byte) error {

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		return err
	}
	defer fmt.Println("close DB connection")
	defer conn.Close(context.Background())

	fmt.Printf("Connect to DB: %s\n", dbname)
	_, err = conn.Exec(context.Background(), "INSERT INTO data (id, val) VALUES ($1, $2);", key, val)
	fmt.Printf("data with key %v added to DB\n", key)
	return err
}
