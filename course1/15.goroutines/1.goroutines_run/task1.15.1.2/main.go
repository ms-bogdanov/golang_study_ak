package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	data := NotifyEvery(ticker, 5*time.Second, "Таймер сработал")
	for v := range data {
		fmt.Println(v)
	}
	fmt.Println("Программа завершена")
}

func NotifyEvery(ticker *time.Ticker, d time.Duration, message string) <-chan string {
	ch := make(chan string)
	timer := time.NewTimer(d)

	go func() {
		defer close(ch)
		for {
			select {
			case <-ticker.C:
				ch <- message
			case <-timer.C:
				ticker.Stop()
				return
			}
		}
	}()

	return ch
}
