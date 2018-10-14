package main

import (
	"admin-go/handlers"
	"admin-go/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {

	gin.SetMode(gin.DebugMode)

	file, _ := os.Create("blog.log")
	gin.DefaultWriter = io.MultiWriter(file)

	router := gin.Default()

	router.POST("/login", handlers.LoginHandler)

	authorized := router.Group("/blog")
	authorized.Use(middleware.JWTAuth())
	{
		authorized.GET("/menu", handlers.MenuHandler)
	}

	router.Run(":8888")
}
