package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

var Db *sql.DB

const (
	getUserById   = "SELECT * FROM users WHERE id = $1"
	createNewUser = "INSERT INTO users VALUES (DEFAULT, $1) RETURNING id"
)

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
}

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres password=admin host=localhost port=5432 dbname=gin_sample sslmode=disable")
	if err != nil {
		os.Exit(2)
	}
}

func (user *Users) GetByID() error {
	err := Db.QueryRow(getUserById, user.ID).Scan(&user.ID, &user.Username)

	return err
}

func (user *Users) CreateNewUser() error {
	err := Db.QueryRow(createNewUser, user.Username).Scan(&user.ID)

	return err
}
