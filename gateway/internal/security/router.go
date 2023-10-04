package security

import (
	"log"

	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterSecurityRoutes(r *gin.Engine, cfg *config.Config) {
	svc := InitAuthServiceClient(cfg)
	securityrepo, err := NewSecurityRepository(cfg, svc)
	securitysvc, err := NewSecurityService(securityrepo)

	if err != nil {
		log.Fatal("Failed to connect to security service grpc: ", err)
	}
	userctrl := NewSecurityController(securitysvc)

	authenticationMiddleware := middleware.NewAuthenticationMiddleware(svc)
	r.POST("/login", userctrl.Login)
	r.GET("/refresh", authenticationMiddleware.Auth(), userctrl.Refresh)
	r.GET("/sendOTP", userctrl.SendOTP)
	r.POST("/validateOTP", userctrl.ValidateOTP)
}
