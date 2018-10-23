package groups

import (
	"admin-go/apis"
	"admin-go/middleware"
	"github.com/gin-gonic/gin"
)

func Video(video *gin.RouterGroup) {

	video.Use(middleware.Limiter(middleware.NewLimit(2)))
	{
		video.GET("/name", apis.VideoApi)
	}
}
