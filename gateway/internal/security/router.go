package security

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
)

func RegisterSecurityRoutes(r *gin.Engine, cfg *config.Config) {
	svc, err := InitAuthServiceClient(cfg)
	if err != nil {
		slog.Fatalf("Failed to connect to security service grpc: %v", err)
	}
	securityrepo := NewSecurityRepository(cfg, svc)
	securitysvc := NewSecurityService(securityrepo)

	userctrl := NewSecurityController(securitysvc)

	authenticationMiddleware := middleware.NewAuthenticationMiddleware(svc)
	r.POST("/login", userctrl.Login)
	r.GET("/refresh", authenticationMiddleware.Auth(), userctrl.Refresh)
	r.POST("/send-otp", userctrl.SendOTP)
	r.GET("/validate-otp", userctrl.ValidateOTP)
}
