package main

import (
	"bitbucket.org/s-kurokawa/gin-sample/src/controllers"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("../static/templates/*.html")

	//router.GET("/", getRoot)
	router.GET("/:id", getUserById)

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
		c.JSON(400, gin.H{"err": err})
		return
	}
	if id <= 0 {
		c.JSON(400, gin.H{"status": "id should be bigger than 0"})
		return
	}

	result := controllers.NewUser().GetByID(id)
	if result == nil || reflect.ValueOf(result).IsNil() {
		c.JSON(404, gin.H{"a": "i"})
		return
	}

	c.JSON(200, result)
}

func config(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}
