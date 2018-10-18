package apis

import (
	"admin-go/handlers"
	"admin-go/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Apis() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	file, _ := os.Create("admin-go.log")
	gin.DefaultWriter = io.MultiWriter(file)

	router := gin.Default()
	router.Use(middleware.CORS())

	router.POST("/login", handlers.LoginHandler)

	authorized := router.Group("/blog")
	authorized.Use(middleware.JWTAuth())
	{
		authorized.GET("/menu", handlers.MenuHandler)
		authorized.GET("/doc", handlers.DocHandler)

	}

	limit := middleware.NewLimit(4)

	player := router.Group("/player")
	player.Use(middleware.Limiter(limit))
	{
		player.GET("/name", handlers.VideoHandler)
	}

	return router
}
