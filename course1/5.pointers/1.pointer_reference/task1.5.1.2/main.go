package main

import (
	"fmt"
)

func main() {
	a := 9900
	mutate(&a)

	s := "macbook"
	ReverseString(&s)

	fmt.Println(a)
	fmt.Println(s)
}

func mutate(a *int) {
	*a = 42
}

func ReverseString(str *string) {

	var res string

	for i := len(*str) - 1; i >= 0; i-- {
		res += string((*str)[i])
	}

	*str = res
}
