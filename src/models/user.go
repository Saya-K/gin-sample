package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

type Users struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"name" xorm:"'nickname'"`
}

type UserRepository struct {
}

func init() {
	var err error
	engine, err = xorm.NewEngine("postgres", "user=postgres password=admin host=localhost port=5432 dbname=gin_sample sslmode=disable")

	if err != nil {
		panic(err)
	}
}

func NewUser(id int, username string) Users {
	return Users{
		ID:       id,
		Username: username,
	}
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (ur UserRepository) GetByID(id int) *Users {
	var user = Users{ID: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}

	return nil
}
