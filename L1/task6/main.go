package main

import (
	"fmt"
	"sync"
	"time"
)

func f1(c <-chan int, q <-chan struct{}, w *sync.WaitGroup) {
	defer w.Done()
	defer fmt.Println("f1 close")
	for {
		select {
		case <-c:
		case <-q:
			return
		}
	}
}

func f2(c <-chan int, w *sync.WaitGroup) {
	defer fmt.Println("f2 close")
	defer w.Done()
	for {
		_, ok := <-c
		if !ok {
			return
		}

	}
}

func f3(c <-chan int, w *sync.WaitGroup) { //??
	defer w.Done()
	defer fmt.Println("f3 close")
	for range c {
		time.Sleep(time.Microsecond)
	}

}

func f4(c <-chan int, d time.Duration, w *sync.WaitGroup) {
	defer w.Done()
	defer fmt.Println("f4 close")
	for {
		select {
		case <-c:
		case <-time.After(d):
			return
		}
	}
}
func f5(c <-chan int, q <-chan struct{}, w *sync.WaitGroup) {
	defer w.Done()
	defer fmt.Println("f5 close")
	<-q
}
func main() {
	var wg sync.WaitGroup
	c1 := make(chan int)
	c2 := make(chan int, 1)
	q := make(chan struct{}, 1)
	wg.Add(5)

	go f4(c1, time.Microsecond, &wg)
	go f2(c2, &wg)
	go f3(c2, &wg)
	go f1(c1, q, &wg)
	go f5(c1, q, &wg)
	q <- struct{}{}
	q <- struct{}{}
	close(c2)
	close(c1)
	wg.Wait()
}
