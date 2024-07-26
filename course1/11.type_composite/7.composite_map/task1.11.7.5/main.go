package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}

	fmt.Println(filterSentence(sentence, filter))
}

func filterSentence(sentence string, filter map[string]bool) string {
	words := strings.Fields(sentence)
	res := []string{}

	for _, w := range words {
		if _, ok := filter[w]; !ok {
			res = append(res, w)
		}
	}
	return strings.Join(res, " ")
}
