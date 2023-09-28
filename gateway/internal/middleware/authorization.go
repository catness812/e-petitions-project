package middleware

import (
	"context"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"path/filepath"
)

type role struct {
	Code        string `yaml:"code"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type resource struct {
	Code        string `yaml:"code"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type permissions struct {
	Role     string `yaml:"role"`
	Resource string `yaml:"resource"`
	Allow    struct {
		Read   bool `yaml:"read"`
		Write  bool `yaml:"write"`
		Delete bool `yaml:"delete"`
	} `yaml:"allow"`
	Deny struct {
		Read   bool `yaml:"read"`
		Write  bool `yaml:"write"`
		Delete bool `yaml:"delete"`
	} `yaml:"deny"`
}

type permissionsConfig struct {
	Roles       []role        `yaml:"roles"`
	Resources   []resource    `yaml:"resources"`
	Permissions []permissions `yaml:"permissions"`
}

func loadConfigRBAC() *permissionsConfig {
	var permConfig *permissionsConfig
	wd, err := os.Getwd()
	if err != nil {
		slog.Fatalf("Failed to get working directory: %v", err)
	}
	configPath := filepath.Join(wd, "rbac.yml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		slog.Fatalf("Failed to read permissions configuration file: %v", err)
	}
	if err := yaml.Unmarshal(data, &permConfig); err != nil {
		slog.Fatalf("Failed to unmarshal YAML data to config: %v", err)
	}
	return permConfig
}

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
