package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

func RateLimiterMiddleware() gin.HandlerFunc {
	limiter := rate.NewLimiter(2, 5)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"status": "Request Failed",
				"error":  "The API is at capacity, try again later.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
