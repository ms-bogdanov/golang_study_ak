package main

import (
	"fmt"
	"time"
)

func main() {

	data := generateData(10)
	go func() {
		time.Sleep(1 * time.Second)
		close(data)
	}()
	for num := range data {
		fmt.Println(num)
	}
}

func generateData(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
