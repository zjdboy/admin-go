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
		log.Fatalf("Reached the limit connections")
		return false
	}

	limiter.Bucket <- 1
	return true
}

func (limiter *Limit) Release() {

	connection := <-limiter.Bucket

	log.Fatalf("Release one connection: %d", connection)
}

func NewLimit(cc int) *Limit {
	return &Limit{
		Concurrent: cc,
		Bucket:     make(chan int, cc),
	}
}

func Limiter(limiter *Limit) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.GetLimiter() {
			c.Abort()
			return
		}

		//defer limiter.Release()
		c.Next()
	}
}
