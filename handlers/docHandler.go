package handlers

import (
	"admin-go/dbops/doc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DocHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"docs": doc.ListData(),
	})
}
