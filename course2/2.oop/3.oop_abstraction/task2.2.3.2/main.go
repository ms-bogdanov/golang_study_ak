package main

import (
	"fmt"
)

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

func (u *User) TableName() string {
	return "users"
}

type Tabler interface {
	TableName() string
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

type FakeDataGenerator interface {
	GenerateFakeUser() User
}

type SQLiteGenerator struct{}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	panic("implement me")
}

func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	panic("implement me")
}

type GoFakeitGenerator struct{}

func (g *GoFakeitGenerator) GenerateFakeUser() User {
	panic("implement me")
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}
	user := User{}

	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)

	for i := 0; i < 10; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}
}
