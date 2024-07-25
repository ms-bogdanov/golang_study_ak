package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	val, newSlice := Pop(a)
	fmt.Printf("Значение: %d, Новый срез: %v", val, newSlice)
}

func Pop(xs []int) (int, []int) {
	val := xs[0]
	newSlice := xs[1:]
	return val, newSlice
}
