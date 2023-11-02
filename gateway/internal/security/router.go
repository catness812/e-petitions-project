package security

import (
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	"github.com/catness812/e-petitions-project/gateway/internal/security/pb"
	"github.com/gofiber/fiber/v2"
)

func RegisterSecurityRoutes(r *fiber.App, securityCtrl *SecurityController, securityClient pb.SecurityServiceClient) {
	authenticationMiddleware := middleware.NewAuthenticationMiddleware(securityClient)
	r.Post("/login", securityCtrl.Login)
	r.Get("/refresh", authenticationMiddleware.Auth(), securityCtrl.Refresh)
	r.Post("/send-otp", securityCtrl.SendOTP)
	r.Get("/validate-otp", securityCtrl.ValidateOTP)
}
