package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64

	fmt.Println(Sin(x))
	fmt.Println(Cos(x))
}

func Sin(x float64) float64 {

	return math.Sin(x)
}

func Cos(x float64) float64 {

	return math.Cos(x)
}
