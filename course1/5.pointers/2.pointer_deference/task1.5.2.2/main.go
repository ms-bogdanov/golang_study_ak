package main

import "fmt"

func main() {
	var a int = 5
	fmt.Println(Factorial(&a))
	var b string = "assa"
	fmt.Println(isPalindrome(&b))
	var numbers []int = []int{30, 1, 2, 5, 1, 30, 29, 30}
	var target int = 30
	fmt.Println(CountOccurrences(&numbers, &target))
	var s string = "golang"
	fmt.Println(ReverseString(&s))
}

func Factorial(n *int) int {
	f := 1
	for i := 1; i <= *n; i++ {
		f *= i
	}
	return f
}

func isPalindrome(str *string) bool {

	if str == nil {
		return false
	}

	var result bool
	var res string

	for _, v := range *str {
		res += string(v)
	}

	if *str == res {
		result = true
	} else {
		result = false
	}

	return result
}

func CountOccurrences(numbers *[]int, target *int) int {

	count := 0
	for _, v := range *numbers {
		if v == *target {
			count++
		}
	}
	return count
}

func ReverseString(str *string) string {

	var res string

	for i := len(*str) - 1; i >= 0; i-- {
		res += string((*str)[i])
	}

	return res
}
