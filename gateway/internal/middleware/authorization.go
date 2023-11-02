package middleware

import (
	"context"
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthMiddleware struct {
	userClient pb.UserServiceClient
	cfg        *config.PermissionsConfig
}

func NewAuthorizationMiddleware(userClient pb.UserServiceClient, rbacCfg *config.PermissionsConfig) *AuthMiddleware {
	return &AuthMiddleware{userClient: userClient, cfg: rbacCfg}
}

func (auth *AuthMiddleware) Authorize(action, resourceCode string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		mail, ok := c.Locals("userEmail").(string)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": "You do not have permission to perform this action.",
			})
		}

		user, err := auth.userClient.GetUserByEmail(context.Background(), &pb.GetUserByEmailRequest{
			Email: mail,
		})
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": "You do not have permission to perform this action.",
			})
		}

		authorized := false
		for _, perm := range auth.cfg.Permissions {
			if perm.Resource == resourceCode && perm.Role == user.Role {
				switch action {
				case "read":
					if perm.Allow.Read {
						authorized = true
					}
				case "write":
					if perm.Allow.Write {
						authorized = true
					}
				case "delete":
					if perm.Allow.Delete {
						authorized = true
					}
				case "update":
					if perm.Allow.Update {
						authorized = true
					}
				}
			}
		}

		if !authorized {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{
				"error":   "Forbidden",
				"message": "You do not have permission to perform this action.",
			})
		}

		return c.Next()
	}
}
