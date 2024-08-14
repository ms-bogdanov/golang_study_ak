package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5}
	result := InsertAfterIDX(xs, 2, 6, 7, 8)
	fmt.Println(result)
}

func InsertAfterIDX(xs []int, idx int, x ...int) []int {
	xs2 := append(xs[:idx+1], x...)
	xs3 := append(xs2, xs[idx+1:]...)
	return xs3
}
