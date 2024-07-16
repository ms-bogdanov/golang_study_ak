package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Длина гипотенузы:", hypotenuse(3.1, 4.325))
}

func hypotenuse(a, b float64) float64 {

	return math.Sqrt(a*a + b*b)
}
