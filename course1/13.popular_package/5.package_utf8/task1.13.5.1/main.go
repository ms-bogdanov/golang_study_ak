package main

import "fmt"

func main() {
	fmt.Println(countUniqueUTF8Chars("Hello, !"))
}

func countUniqueUTF8Chars(s string) int {
	runes := []rune(s)
	uniqueUTF8 := make(map[rune]struct{})
	res := 0

	for _, r := range runes {
		if _, ok := uniqueUTF8[r]; !ok {
			uniqueUTF8[r] = struct{}{}
			res++
		}
	}
	return res
}
