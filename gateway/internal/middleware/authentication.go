package middleware

import (
	"github.com/catness812/e-petitions-project/gateway/internal/security/pb"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationMiddleware struct {
	securityClient pb.SecurityServiceClient
}

func NewAuthenticationMiddleware(securityClient pb.SecurityServiceClient) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{securityClient: securityClient}
}

func (auth *AuthenticationMiddleware) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Request does not contain an access token"})
		}

		token := &pb.Token{Token: tokenString}

		response, err := auth.securityClient.ValidateToken(c.Context(), token)

		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": err.Error()})
		}

		c.Locals("userEmail", response.Email)
		return c.Next()
	}
}
