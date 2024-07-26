package main

import "fmt"

type User struct {
	Nickname string
	Age      int
	Email    string
}

func main() {
	users := []User{User{"max", 27, "qwerty"},
		User{"dima", 27, "asdfg"},
		User{"dima", 27, "asdfg"}}

	uniqueUsers := getUniqueUsers(users)
	fmt.Println(uniqueUsers)
}

func getUniqueUsers(users []User) []User {
	uniqueUsers := make([]User, 0, len(users))
	nick := make(map[string]struct{})
	for _, user := range users {
		if _, ok := nick[user.Nickname]; !ok {
			nick[user.Nickname] = struct{}{}
			uniqueUsers = append(uniqueUsers, user)
		}
	}
	return uniqueUsers
}
