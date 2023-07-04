package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	n = 3
)

func main() {
	//используем небуфферезированный канал
	//горунтина, отправляющая значения в канал, будет заблокирована, пока в канале находится предыдущеее значение
	//получаетель будет заблокирован, пока в канале не появится новое значение
	c := make(chan int)
	go func() {
		timer := time.After(n * time.Second) //канал, по которому отправится значение времени через n * time.Second
		for {
			select {
			case c <- rand.Int():
			case <-timer: // при получении значения из канала timer закроем канал c и завершим горунтину
				close(c)
				return
			}

		}
	}()

	for {
		val, ok := <-c //ok == false когда канал с будет закрыт
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
