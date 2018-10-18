package groups

import (
	"admin-go/apis"
	"admin-go/middleware"
	"github.com/gin-gonic/gin"
)

func Blog(blog *gin.RouterGroup)  {

	blog.Use(middleware.JWTAuth())

	{
		blog.GET("/menu", apis.MenuApi)
		blog.GET("/doc", apis.DocApi)

	}
}
