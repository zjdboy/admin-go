package handlers

import (
	"admin-go/apis"
	"admin-go/config"
	"admin-go/groups"
	"admin-go/middleware"
	"github.com/gin-gonic/gin"
)

func Handlers() *gin.Engine {

	config.Config()

	router := gin.Default()
	router.Use(middleware.CORS())

	router.POST("/login", apis.LoginApi)

	blog := router.Group("/blog")
	groups.Blog(blog)

	video := router.Group("/video")
	groups.Video(video)

	return router
}
