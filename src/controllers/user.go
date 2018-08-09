package controllers

import (
	"bitbucket.org/s-kurokawa/gin-sample/src/models"
)

type User struct {
}

func NewUser() User {
	return User{}
}

func (u User) GetByID(id int) interface{} {
	return models.NewUserRepository().GetByID(id)
}
