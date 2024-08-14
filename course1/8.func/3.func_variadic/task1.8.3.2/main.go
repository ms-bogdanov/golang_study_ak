package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ConcatenateStrings("-", "hello", "world", "how", "are", "you"))
}

func ConcatenateStrings(sep string, str ...string) string {
	var even []string
	var odd []string
	for i, str := range str {
		if i%2 == 0 {
			even = append(even, str)
			continue
		}
		odd = append(odd, str)
	}
	return fmt.Sprintf("even: %s, odd: %s", strings.Join(even, sep), strings.Join(odd, sep))
}
