package main

import "fmt"

var a int = 10
var b int = 3

func main() {
	var sum int
	var difference int
	var product int
	var quotient int
	var remainder int

	sum, difference, product, quotient, remainder = calculate(a, b)
	fmt.Printf("a = %d b = %d sum = %d difference = %d product = %d quotient = %d remainder = %d", a, b, sum, difference,
		product, quotient, remainder)
}

func calculate(a int, b int) (int, int, int, int, int) {
	return a + b, a - b, a * b, a / b, a % b
}
