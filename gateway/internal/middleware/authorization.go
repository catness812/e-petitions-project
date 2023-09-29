package middleware

import (
	"context"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"net/http"
)

type AuthMiddleware struct {
	userClient pb.UserControllerClient
}

func NewAuthMiddleware(userClient pb.UserControllerClient) *AuthMiddleware {
	return &AuthMiddleware{userClient: userClient}
}

func (auth *AuthMiddleware) Authorize(action, resourceCode string) gin.HandlerFunc {
	cfg := loadConfigRBAC()
	return func(c *gin.Context) {
		//mail, ok := c.Get("userMail")
		//if !ok {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		//		"error":   "Unauthorized",
		//		"message": "You do not have permission to perform this action.",
		//	})
		//	return
		//}
		email := "example@email.com"
		//email, ok := mail.(string)
		//if !ok {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		//		"error":   "Unauthorized",
		//		"message": "You do not have permission to perform this action.",
		//	})
		//	return
		//}
		user, err := auth.userClient.GetUserByEmail(context.Background(), &pb.GetUserByEmailRequest{
			Email: email,
		})
		slog.Info(user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You do not have permission to perform this action.",
			})
			return
		}
		authorized := false

		for _, perm := range cfg.Permissions {
			if perm.Resource == resourceCode && perm.Role == user.Role {
				if (action == "read" && perm.Allow.Read) ||
					(action == "write" && perm.Allow.Write) ||
					(action == "delete" && perm.Allow.Delete) {
					authorized = true
					break
				}
			}
		}

		if !authorized {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Forbidden",
				"message": "You do not have permission to perform this action.",
			})
			return
		}
		c.Next()
	}

}
