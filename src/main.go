package main

import (
	"bitbucket.org/s-kurokawa/gin-sample/src/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	router := gin.Default()

	router.GET("/:id", getUserById)
	router.OPTIONS("/createNewUser", options)
	router.POST("/createNewUser", createNewUser)

	router.Run(":8000")
}

//func getRoot(c *gin.Context) {
//
//	c.HTML(200, "main.html", gin.H{})
//}

func getUserById(c *gin.Context) {
	config(c)

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
	config(c)

	user, err := controllers.CreateNewUser(c)
	if err != nil {
		c.JSON(404, gin.H{"status": fmt.Sprint(err)})
		return
	}

	fmt.Println(user)
	c.JSON(200, user)
}

func options(c *gin.Context) {
	config(c)
}
func config(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", c.Request.Header.Get("Access-Control-Request-Headers"))
}
