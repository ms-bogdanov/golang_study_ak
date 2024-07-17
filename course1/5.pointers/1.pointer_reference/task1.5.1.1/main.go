package main

import "fmt"

func main() {
	fmt.Println(*Add(2, 7))
	fmt.Println(*Max([]int{1, 2, 5, 3, 4}))
	fmt.Println(isPrime(7))
	fmt.Println(*ConcatenateStrings([]string{"golang", "_", "study"}))
}

func Add(a, b int) *int {
	sum := a + b
	return &sum
}

func Max(numbers []int) *int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return &max
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func ConcatenateStrings(strs []string) *string {
	var res string = ""
	for _, str := range strs {
		res += str
	}
	return &res
}
