package main

import "fmt"

func main() {
	var name string
	var age int
	var city string

	fmt.Print("Введите Ваше имя:")
	fmt.Scanln(&name)

	fmt.Print("Введите Ваш возраст:")
	fmt.Scanln(&age)

	fmt.Print("Введите Ваш город:")
	fmt.Scanln(&city)

	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Город:", city)

}
