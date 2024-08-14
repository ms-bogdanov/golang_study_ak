package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := generateChan(3)
	ch2 := generateChan(5)
	ch3 := generateChan(7)

	mergeTo := mergeChan2(ch1, ch2, ch3)

	for data := range mergeTo {
		fmt.Println(data)
	}
}

func mergeChan(mergeTo chan int, from ...chan int) {
	defer close(mergeTo)
	var wg sync.WaitGroup

	for _, ch := range from {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for v := range ch {
				mergeTo <- v
			}
		}(ch)
	}
	wg.Wait()
}

func mergeChan2(chans ...chan int) chan int {
	var wg sync.WaitGroup
	resChan := make(chan int, 10)

	for _, ch := range chans {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for v := range c {
				resChan <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(resChan)
	}()
	return resChan
}

func generateChan(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}
