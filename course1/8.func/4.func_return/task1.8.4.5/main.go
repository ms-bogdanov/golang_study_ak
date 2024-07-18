package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CalculatePercentageChange(80.7, 150.5))
}

func CalculatePercentageChange(initialValue, finalValue float64) float64 {

	percentageChange := ((finalValue - initialValue) / initialValue) * 100

	return math.Round(percentageChange*100) / 100
}
