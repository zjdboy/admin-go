package groups

import (
	"admin-go/apis"
	"admin-go/middleware"
	"github.com/gin-gonic/gin"
)

func Video(video *gin.RouterGroup) {
	limit := middleware.NewLimit(100)

	video.Use(middleware.Limiter(limit))
	{
		video.GET("/name", apis.VideoApi)
	}
}
