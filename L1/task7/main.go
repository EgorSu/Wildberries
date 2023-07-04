package main

import (
	"sync"
)

type myMap struct {
	mu sync.Mutex
	ma map[interface{}]interface{}
}

func (m *myMap) set(key, val interface{}) {
	m.mu.Lock()         //блокируем доступ к map перед записью
	defer m.mu.Unlock() //разблокируем доступ к map после записи
	m.ma[key] = val
}

func main() {
	m := &myMap{ma: make(map[interface{}]interface{})}
	//ожидание завершения горутин делаем также как раньше через waitgroup
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				m.set(j, j)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
