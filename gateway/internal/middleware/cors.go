package middleware

import "github.com/gofiber/fiber/v2"

func CorsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Origin, X-Requested-With, Accept, Authorization, Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Allow-Origin")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Max-Age", "3600")

		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusNoContent)
		} else {
			return c.Next()
		}
	}
}
