package main

import (
	"Image-Processing-Service/controllers"
	"Image-Processing-Service/database"

	"github.com/gin-gonic/gin"
)

func main() {
	//database connection
	database.ConnectMongoDB()

	//routes management
	routes := gin.Default()
	routes.POST("/register", controllers.Register)
	routes.POST("/login", controllers.Login)

	routes.Run(":80")
}
