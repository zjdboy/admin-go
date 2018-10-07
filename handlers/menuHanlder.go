package handlers

import (
	"admin-go/dbops/menu"
	"github.com/gin-gonic/gin"
)

func MenuHandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"menus": menu.ListData(),
	})
}
