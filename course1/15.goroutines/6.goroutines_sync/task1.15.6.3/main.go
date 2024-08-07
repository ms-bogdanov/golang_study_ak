package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count atomic.Int64
}

func (c *Counter) Increment() {
	c.count.Add(1)
}

func (c *Counter) GetCount() int64 {
	return c.count.Load()
}

func main() {
	wg := sync.WaitGroup{}
	counter := Counter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println(counter.GetCount())
}
