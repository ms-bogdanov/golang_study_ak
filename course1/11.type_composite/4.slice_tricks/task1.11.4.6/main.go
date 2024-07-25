package main

import "fmt"

func main() {
	xs := []int{1, 2, 3}
	res := InsertToStart(xs, 4, 5, 6)
	fmt.Println(res)
}

func InsertToStart(xs []int, x ...int) []int {
	var newSlice []int
	newSlice = append(newSlice, x...)
	newSlice = append(newSlice, xs...)
	return newSlice
}
