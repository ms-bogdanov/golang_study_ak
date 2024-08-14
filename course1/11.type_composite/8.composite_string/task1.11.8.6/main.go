package main

import (
	"fmt"
)

func main() {
	str := CountVowels("Hello, world!")
	fmt.Println(str)
}

func CountVowels(str string) int {
	count := 0

	for _, value := range str {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}
