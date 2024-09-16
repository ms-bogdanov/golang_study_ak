package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/brianvoe/gofakeit"
	_ "github.com/mattn/go-sqlite3"
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

type SQLiteGenerator struct{}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	t := reflect.TypeOf(table).Elem()
	var columns []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		column := fmt.Sprintf("%s %s", field.Tag.Get("db_field"), field.Tag.Get("db_type"))
		columns = append(columns, column)
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", table.TableName(), strings.Join(columns, ", "))
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

type Migrator struct {
	db           *sql.DB
	sqlGenerator SQLGenerator
}

func NewMigrator(db *sql.DB, sqlGenerator SQLGenerator) *Migrator {
	return &Migrator{db: db, sqlGenerator: sqlGenerator}
}

func (m *Migrator) Migrate(models ...Tabler) error {
	for _, model := range models {
		query := m.sqlGenerator.CreateTableSQL(model)
		_, err := m.db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to create table for model %v: %v", model.TableName(), err)
		}
	}
	return nil
}

type FakeDataGenerator interface {
	GenerateFakeUser() User
}

type GoFakeitGenerator struct{}

func (g *GoFakeitGenerator) GenerateFakeUser() User {
	return User{
		ID:        gofakeit.Number(1, 10000),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
}

func main() {
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	sqlGenerator := &SQLiteGenerator{}
	migrator := NewMigrator(db, sqlGenerator)

	if err := migrator.Migrate(&User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	fmt.Println("Migration completed successfully")

	fakeDataGenerator := &GoFakeitGenerator{}
	for i := 0; i < 10; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}
}
