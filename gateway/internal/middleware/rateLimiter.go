package middleware

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
	"net/http"
)

func RateLimiterMiddleware() fiber.Handler {
	limiter := rate.NewLimiter(2, 5)
	return func(c *fiber.Ctx) error {
		if !limiter.Allow() {
			return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
				"error": "The API is at capacity, try again later.",
			})
		}
		return c.Next()
	}
}
