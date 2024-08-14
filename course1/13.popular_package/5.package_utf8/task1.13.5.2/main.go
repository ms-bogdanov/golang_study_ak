package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Привет, мир"
	res := countRussianLetters(str)
	for k, r := range res {
		fmt.Println(string(k), r)
	}

}

func countRussianLetters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range strings.ToLower(s) {
		if isRussianLetter(r) {
			counts[r]++
		}
	}
	return counts
}

func isRussianLetter(r rune) bool {
	return r >= 'а' && r <= 'я'
}
