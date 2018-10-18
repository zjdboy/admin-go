package apis

import (
	"admin-go/defs"
	"admin-go/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginApi(c *gin.Context) {
	user := &defs.UserCredentials{Username: c.Query("username"), Password: c.Query("password")}

	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Missing Parameter",
		})
		c.Abort()
		return
	}

	token, err := utils.GenerateToken(user.Username)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":http.StatusInternalServerError,
			"msg":"Generate Token Failure",
		})
		log.Fatalln(err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":"Generate Token Successfully",
		"token": token,
	})
}
