package security

import (
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	"github.com/catness812/e-petitions-project/gateway/internal/security/pb"
	"github.com/gin-gonic/gin"
)

func RegisterSecurityRoutes(r *gin.Engine, securityCtrl *SecurityController, securityClient pb.SecurityServiceClient) {
	authenticationMiddleware := middleware.NewAuthenticationMiddleware(securityClient)
	r.POST("/login", securityCtrl.Login)
	r.GET("/refresh", authenticationMiddleware.Auth(), securityCtrl.Refresh)
	r.POST("/send-otp", securityCtrl.SendOTP)
	r.GET("/validate-otp", securityCtrl.ValidateOTP)
}
