package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(concurrentSafeCounter())
}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() int {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
	return c.value
}

func concurrentSafeCounter() int {
	counter := Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	return counter.value
}
