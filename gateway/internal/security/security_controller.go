package security

import (
	"context"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/slog"
)

type ISecurityService interface {
	Login(loginUser model.UserCredentials) (model.Tokens, error)
	Refresh(token string) (model.Tokens, error)
	SendOTP(email string) (string, error)
	ValidateOTP(otp, mail string) (bool, error)
}

type SecurityController struct {
	service    ISecurityService
	userClient pb.UserServiceClient
}

func NewSecurityController(service ISecurityService, userClient pb.UserServiceClient) *SecurityController {
	return &SecurityController{service: service, userClient: userClient}
}

func (ctrl *SecurityController) Login(ctx *fiber.Ctx) error {
	var user model.UserCredentials
	err := ctx.BodyParser(&user)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Invalid request format"})
	}
	tokens, err := ctrl.service.Login(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Invalid credentials"})
	}

	slog.Info("Login request successful")
	return ctx.JSON(fiber.Map{
		"message":       "User successfully logged in",
		"access-token":  tokens.AccessToken,
		"refresh-token": tokens.RefreshToken,
		"userId":        tokens.UserId,
	})
}

func (ctrl *SecurityController) Refresh(ctx *fiber.Ctx) error {
	type refreshToken struct {
		Token string `json:"refreshToken"`
	}
	var rt refreshToken
	err := ctx.BodyParser(&rt)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request format",
		})
	}
	tokens, err := ctrl.service.Refresh(rt.Token)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Could not refresh user session",
		})
	}

	slog.Info("Refresh request successful")
	return ctx.JSON(fiber.Map{
		"message":       "User successfully logged in",
		"access-token":  tokens.AccessToken,
		"refresh-token": tokens.RefreshToken,
	})
}

func (ctrl *SecurityController) SendOTP(ctx *fiber.Ctx) error {
	type otpEmail struct {
		Email string `json:"email"`
	}
	var email otpEmail
	err := ctx.BodyParser(&email)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Invalid request format"})
	}
	_, err = ctrl.service.SendOTP(email.Email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Failed to send OTP"})
	}

	slog.Info("OTP sent successfully")
	return ctx.JSON(fiber.Map{"error": false, "message": "OTP sent successfully"})
}

func (ctrl *SecurityController) ValidateOTP(ctx *fiber.Ctx) error {
	otp := ctx.Query("otp")
	email := ctx.Query("email")
	if otp == "" || email == "" {
		slog.Error("Failed to validate OTP")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Failed to validate OTP"})
	}
	validated, err := ctrl.service.ValidateOTP(otp, email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Failed to validate OTP"})
	}
	_, err = ctrl.userClient.CreateUserOTP(context.Background(), &pb.UserRequest{Email: email, Password: otp})
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Failed to validate OTP"})
	}
	slog.Info("OTP successfully validated")
	return ctx.JSON(fiber.Map{
		"error":     false,
		"message":   "OTP successfully validated",
		"validated": validated,
	})
}
