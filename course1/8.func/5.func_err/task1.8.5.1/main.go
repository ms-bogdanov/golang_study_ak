package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Println(CalculatePercentageChange("34.345", "45.45"))

}

func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {

	parseInitialValue, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		return 0, err
	}

	parseFinalValue, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0, err
	}

	percentageChange := ((parseFinalValue - parseInitialValue) / parseInitialValue) * 100

	return math.Round(percentageChange*100) / 100, err
}
