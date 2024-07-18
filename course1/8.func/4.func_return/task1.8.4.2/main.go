package main

import "fmt"

func main() {
	fmt.Println(FindMaxAndMin(8, 3, 5, 29, 9, 15, 4))
}

func FindMaxAndMin(n ...int) (int, int) {

	var max int = n[0]
	var min int = n[0]

	for _, v := range n {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return max, min
}
