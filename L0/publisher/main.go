package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/nats-io/stan.go"
)

const (
	pathFile  = "data"
	natsTheme = "test"
	natsURL   = stan.DefaultNatsURL
	clusterID = "test-cluster"
	clientID  = "1"
)

type Data struct {
	mu   sync.RWMutex //!!!
	data []*os.File
}

func getFileData(path string, data *Data) error {

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Printf("open file %s\n", file.Name())
		jsonFile, err := os.Open(fmt.Sprintf("%s/%s", pathFile, file.Name()))
		if err != nil {
			return err
		}
		data.data = append(data.data, jsonFile)
	}
	return nil
}

func main() {

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	defer sc.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	data := &Data{data: make([]*os.File, 0)}

	err = getFileData(pathFile, data)
	if err != nil {
		fmt.Printf("error with file: %v\n", err)
		return
	}
	for _, file := range data.data {
		byteValue, _ := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("json unmarshal error: %v\n", err)
			return
		}
		err = sc.Publish(natsTheme, byteValue)
		if err != nil {
			fmt.Printf("Unable to publish to NATS: %v\n", err)
			return
		}
		fmt.Printf("file %v published\n", file.Name())
		//time.Sleep(300 * time.Millisecond)
	}
}
