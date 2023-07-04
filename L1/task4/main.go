package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)     //  канал для сигнала SIGINT
	signal.Notify(sigs, syscall.SIGINT) // перенаправляем сигнал SIGINT в канал sigs

	ch := make(chan interface{}, 3)
	var wg sync.WaitGroup
	//для получения кол-ва воркеров будем исползовать flag
	//flag вернёт указатель на значение
	num := flag.Int("worker", 1, "number of worker")
	flag.Parse() // получаем значения флагов из командной строки
	for i := 0; i < *num; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer fmt.Printf("done worker %d\n", i)
			for val := range ch { // при пустом канале цикл приостановится и будет ждать значения в канале
				fmt.Printf("%v\n", val)
			}
		}()
	}
Loop:
	for {
		//с помощью select будем отправлять случайные данные в канал с
		select {
		case ch <- rand.Intn(100):
		case ch <- rand.Float32():
		case ch <- "go":
		case ch <- struct{}{}:
		case <-sigs:
			close(ch)
			break Loop
		}
	}
	//при получение значения из канала sigs горунитины завершат цикл range (канал ch закрыт)
	//дождёмся завершения горутин
	wg.Wait()

}
