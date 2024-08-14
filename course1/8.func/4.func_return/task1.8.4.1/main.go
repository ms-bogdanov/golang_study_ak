package main

import "fmt"

func main() {
	div, rem := DivideAndRemainder(9, 3)
	fmt.Printf("Частное: %d, Остаток: %d\n", div, rem)
}

func DivideAndRemainder(a, b int) (int, int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Деление на ноль")
		}
	}()

	return a / b, a % b
}
