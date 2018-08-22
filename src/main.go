package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.Default()

	// CORS(Cross-Origin Resource Sharing)対応
	// @see https://github.com/gin-contrib/cors
	// 全ての接続元を許可
	router.Use(cors.Default())

	//router.GET("/:id", getUserById)
	//router.POST("/createNewUser", createNewUser)

	router.POST("/login", login)
	router.POST("/user/:id", getUserInfo)
	router.POST("/newUser", createNewUser)

	router.Run(":8081")
}

// for vue-sample
//func getUserById(c *gin.Context) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		c.JSON(400, gin.H{"status": fmt.Sprint(err)})
//		return
//	}
//	if id <= 0 {
//		c.JSON(400, gin.H{"status": "id should be bigger than 0"})
//		return
//	}
//
//	user, err := controllers.GetByID(id)
//	if err != nil {
//		c.JSON(404, gin.H{"status": fmt.Sprint(err)})
//		return
//	}
//
//	c.JSON(200, user)
//}
//
//func createNewUser(c *gin.Context) {
//	user, err := controllers.CreateNewUser(c)
//	if err != nil {
//		c.JSON(404, gin.H{"status": fmt.Sprint(err)})
//		return
//	}
//
//	c.JSON(200, user)
//}

// Following are for the connection between auth-api and auth-front
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

func getUserInfo(c *gin.Context) {
	type Token struct {
		Token string `json:"token"`
	}

	var token Token
	err := c.BindJSON(&token)
	if err != nil {
		c.JSON(401, gin.H{"result": false, "msg": "failed"})
		return
	}

	if token.Token != "abcdefghijklmnop" {
		c.JSON(401, gin.H{"result": false, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"name": "kurokawa", "email": "kurokawa@psi-net.co.jp", "password": "kurokawa", "loggedInAt": time.Now().Format("2006/01/02 15:04:05"), "createdAt": "2018/08/22 15:52:45"})
	}
}

func createNewUser(c *gin.Context) {
	type User struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(401, gin.H{"result": false, "msg": "failed"})
		return
	}

	if user.Email != "kurokawa@psi-net.co.jp" {
		c.JSON(401, gin.H{"result": false, "msg": "failed"})
	} else {
		c.JSON(200, gin.H{"result": true, "msg": "success"})
	}
}
