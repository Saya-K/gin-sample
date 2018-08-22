package main

import (
	"bitbucket.org/s-kurokawa/gin-sample/src/controllers"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	router := gin.Default()

	// CORS(Cross-Origin Resource Sharing)対応
	// @see https://github.com/gin-contrib/cors
	// 全ての接続元を許可
	router.Use(cors.Default())

	router.GET("/:id", getUserById)
	router.POST("/createNewUser", createNewUser)

	router.POST("/login", login)

	router.Run(":8081")
}

func getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"status": fmt.Sprint(err)})
		return
	}
	if id <= 0 {
		c.JSON(400, gin.H{"status": "id should be bigger than 0"})
		return
	}

	user, err := controllers.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"status": fmt.Sprint(err)})
		return
	}

	c.JSON(200, user)
}

func createNewUser(c *gin.Context) {
	user, err := controllers.CreateNewUser(c)
	if err != nil {
		c.JSON(404, gin.H{"status": fmt.Sprint(err)})
		return
	}

	c.JSON(200, user)
}

func login(c *gin.Context) {
	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(401, gin.H{"result": false, "token": "", "msg": "failed"})
		return
	}

	if user.Password == "a" {
		c.JSON(401, gin.H{"result": false, "token": "", "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"result": true, "token": "abcdefghijklmnop", "msg": "success"})
	}
}
