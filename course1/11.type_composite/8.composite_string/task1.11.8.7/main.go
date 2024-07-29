package main

import "fmt"

func main() {
	result := ReplaceSymbols("Hello, world!", 'o', '0')
	fmt.Println(result)
}

func ReplaceSymbols(str string, old, new rune) string {
	var newString string

	for _, sym := range str {
		if sym == old {
			newString += string(new)
		} else {
			newString += string(sym)
		}
	}
	return newString
}
