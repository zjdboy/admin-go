package config

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Config()  {

	gin.SetMode(gin.DebugMode)

	file, _ := os.Create("admin-go.log")

	gin.DefaultWriter = io.MultiWriter(file)
}
