package handlers

import (
	"github.com/gin-gonic/gin"
)

func DocHandler(c *gin.Context) {
	c.JSON(200, nil)
}
