package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(createUniqueText("bar bar bar foo foo baz"))
}

func createUniqueText(text string) string {
	words := strings.Fields(text)
	myMap := make(map[string]struct{})
	res := []string{}

	for _, w := range words {
		if _, ok := myMap[w]; !ok {
			myMap[w] = struct{}{}
			res = append(res, w)
		}
	}
	return strings.Join(res, " ,")
}
