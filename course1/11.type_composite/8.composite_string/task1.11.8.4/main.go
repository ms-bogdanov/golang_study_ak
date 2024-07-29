package main

import (
	"fmt"
	"strings"
)

func main() {
	result := concatStrings("Hello", " ", "world!")
	fmt.Println(result)
}

func concatStrings(xs ...string) string {
	return strings.Join(xs, "")
}
