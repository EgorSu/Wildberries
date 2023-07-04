package main

import (
	"fmt"
	"sync"
)

// делаем как в задании 2
func main() {

	var wg sync.WaitGroup
	var number []int = []int{2, 4, 6, 8, 10}
	c := make(chan int, len(number))

	for _, val := range number {
		val := val
		wg.Add(1)
		go func(c chan<- int) {
			c <- val * val
			wg.Done()
		}(c)
	}
	wg.Wait()
	close(c)

	var sum int
	for val := range c {
		sum += val
	}
	fmt.Println(sum)
}
