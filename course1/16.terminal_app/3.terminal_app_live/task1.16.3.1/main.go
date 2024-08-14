package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		time.Sleep(time.Second)
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Текущее время: %v\n", time.Now().Format("15:04:05"))
		fmt.Printf("Текущая дата: %v\n", time.Now().Format("2006-01-02"))
	}
}
