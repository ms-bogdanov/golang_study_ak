package main

import "fmt"

func main() {
	var a int = 10
	var s string = "hello"

	fmt.Println(getVariableType(a))
	fmt.Println(getVariableType(s))
}

func getVariableType(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}
