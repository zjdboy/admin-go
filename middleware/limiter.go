package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Limit struct {
	Concurrent int
	Bucket     chan int
}

func (limiter *Limit) GetLimiter() bool {
	if len(limiter.Bucket) >= limiter.Concurrent {
		log.Fatalf("Reached the limt connections")
		return false
	}

	limiter.Bucket <- 1
	return true
}

func (limiter *Limit) Release() {

	connection := <-limiter.Bucket

	log.Fatalf("Release one connection: %d", connection)
}

func Limiter(cc int) gin.HandlerFunc {
	limiter := &Limit{
		Concurrent: cc,
		Bucket:     make(chan int, cc),
	}
	return func(c *gin.Context) {
		if !limiter.GetLimiter() {
			c.Abort()
			return
		}
		c.Next()
	}
}
