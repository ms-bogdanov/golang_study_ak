package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	occurrences := countWordOccurrences(text)

	for word, count := range occurrences {
		fmt.Printf("%s: %d\n", word, count)
	}

}

func countWordOccurrences(text string) map[string]int {
	words := strings.Fields(text)
	countWords := make(map[string]int)

	for _, word := range words {
		countWords[word]++
	}

	return countWords
}
