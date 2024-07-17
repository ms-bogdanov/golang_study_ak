package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(UserInfo("Иван", 25, "Москва", "Санкт-Петербург", "Казань"))
}

func UserInfo(name string, age int, cities ...string) string {
	return fmt.Sprintf("Имя: %s, возраст: %d, города: %s", name, age, strings.Join(cities, ", "))
}
