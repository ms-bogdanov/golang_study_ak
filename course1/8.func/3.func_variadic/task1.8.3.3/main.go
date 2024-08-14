package main

import "fmt"

func main() {
	PrintNumbers(1, 22, 3, 34, 5, 10, 20)
}

func PrintNumbers(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}
