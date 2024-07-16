package main

import "fmt"

func main() {
	var a int = 1010
	var b int = 20

	fmt.Println("a & b =", bitwiseAnd(a, b))

	fmt.Println("a | b =", bitwiseOr(a, b))

	fmt.Println("a ^ b =", bitwiseXor(a, b))

	fmt.Println("a << b =", bitwiseLeftShift(a, b))

	fmt.Println("a >> b =", bitwiseRightShift(a, b))
}

func bitwiseAnd(a, b int) int {
	return a & b
}

func bitwiseOr(a, b int) int {
	return a | b
}

func bitwiseXor(a, b int) int {
	return a ^ b
}

func bitwiseLeftShift(a, b int) int {
	return a << b
}

func bitwiseRightShift(a, b int) int {
	return a >> b
}
