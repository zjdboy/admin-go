package middleware

import (
	"admin-go/defs"
	"admin-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(defs.HEADER_FIELD_SESSION)
		uname := c.GetHeader(defs.HEADER_FIELD_UNAME)

		if token == "" || uname == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "Missing Header",
			})
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "Authentication Failure",
			})
			c.Abort()
			return
		}

		if claims != nil {
			if claims.ExpiresAt < time.Now().Unix() {
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "Login Timeout",
				})
				c.Abort()
				return
			}
			if claims.Username != uname {
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "Bad Request",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
