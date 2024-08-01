package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(GenerateMathString([]int{2, 4, 6}, "+"))
}

func GenerateMathString(operands []int, operator string) string {
	if len(operands) <= 1 {
		return "add a number"
	}

	switch operator {
	case "+":
		return summ(operands)
	case "-":
		return sub(operands)
	case "*":
		return mult(operands)
	case "/":
		return div(operands)
	default:
		return "Unknown operator"
	}
}

func summ(operands []int) string {
	var res int
	var str []string

	for _, operand := range operands {
		res += operand
		str = append(str, fmt.Sprintf("%d", operand))
	}
	return fmt.Sprintf("%s = %d", strings.Join(str, " + "), res)
}

func sub(operands []int) string {
	res := operands[0]
	var str []string
	str = append(str, fmt.Sprintf("%d", operands[0]))

	for i := 1; i < len(operands); i++ {
		res -= operands[i]
		str = append(str, fmt.Sprintf("%d", operands[i]))
	}
	return fmt.Sprintf("%s = %d", strings.Join(str, " - "), res)
}

func mult(operands []int) string {
	var res int = 1
	var str []string

	for _, operand := range operands {
		res *= operand
		str = append(str, fmt.Sprintf("%d", operand))
	}
	return fmt.Sprintf("%s = %d", strings.Join(str, " * "), res)
}

func div(operands []int) string {
	res := float64(operands[0])
	var str []string
	str = append(str, fmt.Sprintf("%d", operands[0]))

	for i := 1; i < len(operands); i++ {
		res /= float64(operands[i])
		str = append(str, fmt.Sprintf("%d", operands[i]))
	}
	return fmt.Sprintf("%s = %.2f", strings.Join(str, " / "), res)
}
