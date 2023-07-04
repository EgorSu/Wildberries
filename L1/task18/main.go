package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// сделаем два счётчика: через mutex и через atomic
type counter struct {
	mu            sync.Mutex
	counter       uint64
	counterAtomic uint64
}

func (c *counter) add() {
	c.mu.Lock()         //блокируем доступ к counter
	defer c.mu.Unlock() //при завершении метода разблокируем доступ к counter
	c.counter++
}

func (c *counter) addAtomic() {
	//атомарный счётчик не нуждается в блокировках доступа
	atomic.AddUint64(&c.counterAtomic, 1)
}

func main() {

	c := counter{counter: 0}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				c.add()
				c.addAtomic()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("mutex lock counter %d\natomic counter %d\n", c.counter, c.counterAtomic)
}
