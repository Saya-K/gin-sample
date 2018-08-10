package controllers

import (
	"bitbucket.org/s-kurokawa/gin-sample/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetByID(id int) (interface{}, error) {
	var user models.Users
	var err error

	user.ID = id

	err = user.GetByID()
	if err != nil {
		return struct{}{}, err
	}
	if user.Username == "" {
		err = fmt.Errorf("no user has id:" + fmt.Sprint(id))
		return struct{}{}, err
	}

	return user, err
}

func CreateNewUser(c *gin.Context) (interface{}, error) {
	var user models.Users
	err := c.BindJSON(&user)
	if err != nil {
		return struct{}{}, err
	}

	err = user.CreateNewUser()
	if err != nil {
		return struct{}{}, err
	}

	return user, err
}
