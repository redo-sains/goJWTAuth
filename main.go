package main

import (
	// "fmt"
	"jwtAuth/controllers"
	"jwtAuth/initializers"
	"jwtAuth/middlewares"

	"github.com/gin-gonic/gin"
	// "os"
)

func init() {
	initializers.DotEnv()
	initializers.InitDB()
	initializers.SyncDatabase()
}

func main() {

	router := gin.Default()

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middlewares.Validate, controllers.GetUser)

	router.Run()
}
