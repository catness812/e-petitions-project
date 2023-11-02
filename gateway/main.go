package main

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	"github.com/catness812/e-petitions-project/gateway/internal/petition"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/catness812/e-petitions-project/gateway/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(middleware.RateLimiterMiddleware())
	r.Use(corsMiddleware())
	registerRoutes(r)
	r.GET("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success",
		})
	})
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
		//origin := c.Request.Header.Get("Origin")
		//if origin == "https://epetitii.co" {
		//	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		//}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}
