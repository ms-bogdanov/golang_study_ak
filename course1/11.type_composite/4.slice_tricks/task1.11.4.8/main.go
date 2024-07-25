package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5}
	firstElement, shiftedSlice := Shift(xs)
	fmt.Println(firstElement)
	fmt.Println(shiftedSlice)
}

func Shift(xs []int) (int, []int) {

	if len(xs) == 0 {
		return 0, xs
	}

	if len(xs) == 1 {
		return xs[0], xs
	}

	xs = append(xs[len(xs)-1:], xs[:len(xs)-1]...)
	return xs[1], xs
}
