package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
)

type User struct {
	name string
	age  int
}

func main() {
	users := getUsers()
	result := preparePrint(users)
	fmt.Println(result)
}

func getUsers() []User {
	users := make([]User, 10)
	for i := 0; i < 10; i++ {
		users[i] = User{
			gofakeit.Name(),
			gofakeit.Number(18, 60),
		}
	}
	return users
}

func preparePrint(users []User) string {
	var str string
	for _, user := range users {
		str += fmt.Sprintf("Имя: %s, Возраст: %d\n", user.name, user.age)
	}
	return str
}
