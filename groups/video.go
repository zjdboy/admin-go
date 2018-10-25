package groups

import (
	"admin-go/apis"
	"admin-go/middleware"
	"github.com/gin-gonic/gin"
)

func Video(video *gin.RouterGroup) {
	limiter := middleware.NewLimit(3)
	video.Use(middleware.Limiter(limiter))
	{
		video.GET("/name", apis.VideoApi)
	}
}
