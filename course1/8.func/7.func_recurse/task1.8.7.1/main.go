package main

import "fmt"

func main() {
	fmt.Println(Factorial(7))
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return Factorial(n-1) * n
}
