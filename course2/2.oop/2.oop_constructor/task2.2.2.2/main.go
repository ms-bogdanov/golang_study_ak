package main

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func WithUsername(Username string) UserOption {
	return func(u *User) {
		u.Username = Username
	}
}

func WithEmail(Email string) UserOption {
	return func(u *User) {
		u.Email = Email
	}
}

func WithRole(Role string) UserOption {
	return func(u *User) {
		u.Role = Role
	}
}

func NewUser(ID int, options ...UserOption) *User {
	users := &User{
		ID: ID,
	}

	for _, option := range options {
		option(users)
	}
	return users
}

func main() {
	user := NewUser(1,
		WithUsername("testuser"),
		WithEmail("testuser@example.com"),
		WithRole("admin"))

	fmt.Printf("User: %+v\n", user)
}
