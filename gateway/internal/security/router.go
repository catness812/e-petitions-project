package security

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	"github.com/catness812/e-petitions-project/gateway/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
)

func RegisterSecurityRoutes(r *gin.Engine, cfg *config.Config) {
	svc, err := InitAuthServiceClient(cfg)

	if err != nil {
		slog.Fatalf("Failed to connect to security service grpc: %v", err)
	}

	userClient := user.InitUserServiceClient(cfg)
	securityrepo := NewSecurityRepository(cfg, svc)
	securitysvc := NewSecurityService(securityrepo)
	userRepo, err := user.NewUserRepository(cfg, userClient)
	if err != nil {
		slog.Fatalf("Failed to initialize user repository for the security gateway")
	}
	userSvc, err := user.NewUserService(userRepo)
	if err != nil {
		slog.Fatalf("Failed to initialize user service for the security gateway")
	}
	securityCtrl := NewSecurityController(securitysvc, userSvc)

	authenticationMiddleware := middleware.NewAuthenticationMiddleware(svc)
	r.POST("/login", securityCtrl.Login)
	r.GET("/refresh", authenticationMiddleware.Auth(), securityCtrl.Refresh)
	r.POST("/send-otp", securityCtrl.SendOTP)
	r.GET("/validate-otp", securityCtrl.ValidateOTP)
}
