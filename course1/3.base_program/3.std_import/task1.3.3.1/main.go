package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	result := Sqrt(x)
	fmt.Println(result)
}

func Sqrt(x float64) float64 {

	return math.Sqrt(x)
}
