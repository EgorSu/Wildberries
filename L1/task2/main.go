package main

import (
	"fmt"
	"sync"
)

func main() {

	var number [5]int = [5]int{2, 4, 6, 8, 10}
	//для передачи данных будем использовать буфферезированный канал
	c := make(chan int, len(number))
	var wg sync.WaitGroup
	for _, val := range number {
		//переменная val является общей для всех горутин
		// создадим новые экземпляры val для каждоой горутины
		val := val
		wg.Add(1) //увеличиваем счётчик горутин
		go func() {
			defer wg.Done() //при заверешении горутины сообщаем об этом wg
			c <- val * val  //отправяем результат операции в канал
		}()
	}
	wg.Wait() //ожидаем окончания всех горутин
	close(c)  //закрываем канал
	for range number {
		fmt.Println(<-c)
	}
}
