package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5}
	result := Cut(xs, 1, 3)
	fmt.Println(result)
}

func Cut(xs []int, start, end int) []int {
	if start < 0 || end < 0 || start >= len(xs) || end >= len(xs) {
		return []int{}
	}
	return xs[start : end+1]
}
