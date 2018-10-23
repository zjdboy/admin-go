package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Limit struct {
	Size   int
	Bucket chan int
}

func (limiter *Limit) GetLimiter() bool {
	if len(limiter.Bucket) >= limiter.Size {
		log.Printf("Reached the limit connections")
		return false
	}

	limiter.Bucket <- 1
	return true
}

func (limiter *Limit) Release() {

	connection := <-limiter.Bucket

	log.Printf("Release one connection: %d", connection)
}

func NewLimit(cc int) *Limit {
	return &Limit{
		Size:   cc,
		Bucket: make(chan int, cc),
	}
}

func Limiter(limiter *Limit) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.GetLimiter() {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "Reached the limit connections",
			})
			c.Abort()
		}
		c.Next()
		defer limiter.Release()
	}
}
