package main

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/catness812/e-petitions-project/gateway/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())
	registerRoutes(r)
	err := r.Run(":1337")
	if err != nil {
		slog.Fatalf("Failed to start server: %v", err)
	}
}

func registerRoutes(r *gin.Engine) {
	cfg := config.LoadConfig()
	rbacCfg := config.LoadConfigRBAC()
	securityClient, err := security.InitAuthServiceClient(cfg)
	if err != nil {
		slog.Fatalf("Failed to connect to security service grpc: %v", err)
	}
	userClient := user.InitUserServiceClient(cfg)
	securityRepo := security.NewSecurityRepository(cfg, securityClient)
	securitySvc := security.NewSecurityService(securityRepo)
	securityCtrl := security.NewSecurityController(securitySvc, userClient)

	userRepo := user.NewUserRepository(cfg, userClient)
	userSvc := user.NewUserService(userRepo)
	userCtrl := user.NewUserController(userSvc)

	user.RegisterUserRoutes(r, rbacCfg, userCtrl, userClient, securityClient)
	petition.RegisterPetitionRoutes(r, cfg)
	security.RegisterSecurityRoutes(r, securityCtrl, securityClient)

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,Access-Control-Allow-Origin")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
