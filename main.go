package main

import (
	"Image-Processing-Service/controllers"
	"Image-Processing-Service/database"
	"Image-Processing-Service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	//database connection
	database.ConnectMongoDB()

	//routes management
	routes := gin.Default()
	routes.MaxMultipartMemory = 8 << 20 // 8 MB
	routes.RedirectTrailingSlash = false

	// Public routes
	routes.POST("/register", controllers.Register)
	routes.POST("/login", controllers.Login)
	routes.Static("/uploads", "./assets/uploads")
	routes.Static("/edited", "./assets/edited")

	//Authorized routes
	authorized := routes.Group("/images")
	authorized.Use(middleware.JwtAuth)
	{
		authorized.POST("", controllers.UploadImage) // /image routes
	}

	routes.Run(":80")
}
