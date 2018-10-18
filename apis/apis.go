package apis

import (
	"admin-go/config"
	"admin-go/handlers"
	"admin-go/middleware"
	"github.com/gin-gonic/gin"
)

func Apis() *gin.Engine {

	config.Config()

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
