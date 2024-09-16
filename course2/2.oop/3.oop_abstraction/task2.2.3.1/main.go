package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/brianvoe/gofakeit"
)

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

type Tabler interface {
	TableName() string
}

func (u *User) TableName() string {
	return "user"
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
	t := reflect.TypeOf(table).Elem()
	var columns []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		column := fmt.Sprintf("%s %s", field.Tag.Get("db_field"), field.Tag.Get("db_type"))
		columns = append(columns, column)
	}
	return fmt.Sprintf("CREATE TABLE %s (%s);", table.TableName(), strings.Join(columns, ", "))
}

func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	v := reflect.ValueOf(model).Elem()
	t := reflect.TypeOf(model).Elem()

	var columns, placeholders []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columns = append(columns, field.Tag.Get("db_field"))
		placeholders = append(placeholders, fmt.Sprintf("'%v'", v.Field(i).Interface()))
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", model.TableName(), strings.Join(columns, ", "), strings.Join(placeholders, ", "))
}

type GoFakeitGenerator struct{}

func (g *GoFakeitGenerator) GenerateFakeUser() User {
	return User{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
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
