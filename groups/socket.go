package groups

import (
	"admin-go/apis"
	"github.com/gin-gonic/gin"
)

func Socket(group *gin.RouterGroup)  {
	group.GET("/server", apis.Server)
	group.GET("/room", apis.Room)
}
