package main

import "fmt"

func main() {
	var a int = 5
	var b float64 = 0.5
	var c string = "Hello, world!"
	var d bool = true

	changeInt(&a)
	changeFloat(&b)
	changeString(&c)
	changeBool(&d)

	fmt.Println(a, b, c, d)
}

func changeInt(a *int) {
	*a = 20
}

func changeFloat(b *float64) {
	*b = 6.28
}

func changeString(c *string) {
	*c = "Goodbye, world!"
}

func changeBool(d *bool) {
	*d = false
}
