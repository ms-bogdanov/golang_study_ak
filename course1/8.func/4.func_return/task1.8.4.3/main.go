package main

import "fmt"

func main() {
	fmt.Println(CalculateStockValue(253.45, 30))
}

func CalculateStockValue(price float64, quantity int) (float64, float64) {
	return price * float64(quantity), price
}
