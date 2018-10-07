package main

import (
	"admin-go/handlers"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {

	gin.SetMode(gin.DebugMode)

	file, _ := os.Create("blog.log")
	gin.DefaultWriter = io.MultiWriter(file)

	router := gin.Default()

	router.GET("/menu", handlers.MenuHandler)

	router.Run(":8888")
}
