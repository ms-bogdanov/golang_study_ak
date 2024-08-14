package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	result := Floor(x)
	fmt.Println(result)
}

func Floor(x float64) float64 {

	return math.Floor(x)
}
