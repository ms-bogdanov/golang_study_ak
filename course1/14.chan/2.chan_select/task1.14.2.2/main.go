package main

import (
	"fmt"
	"time"
)

func main() {
	timeoutFunc := timeout(1 * time.Second)
	since := time.NewTimer(3050 * time.Millisecond)

	for {
		select {
		case <-since.C:
			fmt.Println("Функция не выполнена вовремя")
			return
		default:
			if timeoutFunc() {
				fmt.Println("Функция выполнена вовремя")
				return
			}
		}
	}
}

func timeout(timeout time.Duration) func() bool {
	сh := make(chan bool)

	go func() {
		time.Sleep(timeout)
		сh <- true
		close(сh)
	}()

	return func() bool {
		select {
		case <-сh:
			return false
		default:
			return true
		}
	}
}
