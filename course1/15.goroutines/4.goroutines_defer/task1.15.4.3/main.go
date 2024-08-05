package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	ch := make(chan string)
	myPanic(ch)
	fmt.Println(<-ch)
}

func myPanic(ch chan string) {
	panic("my panic message")
}
