package apis

import (
	"admin-go/dbops/menu"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MenuApi(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"menus": menu.ListData(),
	})
}
