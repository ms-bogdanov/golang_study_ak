package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var DB = &sql.DB{}

func initDbConn() {
	db, err := sql.Open("sqlite3", "./course1/13.popular_package/10.package_db/task1.13.10.1/users.db")
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())
	}
	DB = db
}

func CreateUserTable() error {
	_, err := DB.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
		)`)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(user User) error {
	_, err := DB.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func SelectUser(id int) (User, error) {
	var user User

	err := DB.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func UpdateUser(user User) error {
	_, err := DB.Exec("UPDATE users SET age = ?, name = ? WHERE id = ?", user.Age, user.Name, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	_, err := DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
