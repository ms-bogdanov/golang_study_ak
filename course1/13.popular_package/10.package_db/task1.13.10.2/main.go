package main

import (
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
)

const (
	Insert      = "INSERT"
	Select      = "SELECT"
	Update      = "UPDATE"
	Delete      = "DELETE"
	table       = "users"
	idCol       = "id"
	usernameCol = "username"
	emailCol    = "email"
)

type User struct {
	ID       int
	Username string
	Email    string
}

var DB = &sql.DB{}

func initDbConn() {
	db, err := sql.Open("sqlite3", "./course1/13.popular_package/10.package_db/task1.13.10.2/users.db")
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
	query, args, err := PrepareQuery(Insert, table, user)
	if err != nil {
		return err
	}

	_, err = DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func SelectUser(userID int) (User, error) {
	user := User{
		ID: userID,
	}

	query, args, err := PrepareQuery(Insert, table, user)
	if err != nil {
		return User{}, err
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return User{}, err
	}

	for rows.Next() {
		err = rows.Scan(&user.Username, &user.Email)
		if err != nil {
			return User{}, err
		}
	}

	return user, nil
}

func UpdateUser(user User) error {
	query, args, err := PrepareQuery(Insert, table, user)
	if err != nil {
		return err
	}

	_, err = DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(userID int) error {
	user := User{
		ID: userID,
	}

	query, args, err := PrepareQuery(Insert, table, user)
	if err != nil {
		return err
	}

	_, err = DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func PrepareQuery(operation string, table string, user User) (string, []interface{}, error) {
	switch operation {
	case Insert:
		builder := sq.Insert(table).
			Columns(usernameCol, emailCol).
			Values(user.Username, user.Email)
		return builder.ToSql()

	case Select:
		builder := sq.Select(usernameCol, user.Email).
			From(table).
			Where(sq.Eq{idCol: user.ID})
		return builder.ToSql()

	case Update:
		builder := sq.Update(table).
			SetMap(sq.Eq{usernameCol: user.Username, emailCol: user.Email}).
			Where(sq.Eq{idCol: user.ID})
		return builder.ToSql()

	case Delete:
		builder := sq.Delete(table).
			Where(sq.Eq{idCol: user.ID})
		return builder.ToSql()

	default:
		return "", nil, fmt.Errorf("unknown operation")
	}
}

func main() {
	initDbConn()
	defer DB.Close()
}
