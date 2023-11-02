package security

import (
	"context"
	"net/http"

	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"github.com/gookit/slog"

	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gin-gonic/gin"
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

func (ctrl *SecurityController) Login(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	tokens, err := ctrl.service.Login(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	slog.Info("Login request successful")
	ctx.JSON(http.StatusOK, gin.H{
		"access-token":  tokens.AccessToken,
		"refresh-token": tokens.RefreshToken,
		"userId":        tokens.UserId,
	})
}

func (ctrl *SecurityController) Refresh(ctx *gin.Context) {
	type refreshToken struct {
		Token string `json:"refreshToken"`
	}
	var rt refreshToken
	err := ctx.BindJSON(&rt)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	tokens, err := ctrl.service.Refresh(rt.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	slog.Info("Refresh request successful")
	ctx.JSON(http.StatusOK, gin.H{
		"access-token":  tokens.AccessToken,
		"refresh-token": tokens.RefreshToken,
	})
}

func (ctrl *SecurityController) SendOTP(ctx *gin.Context) {
	type otpEmail struct {
		Email string `json:"email"`
	}
	var email otpEmail
	err := ctx.BindJSON(&email)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = ctrl.service.SendOTP(email.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	slog.Info("OTP sent successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func (ctrl *SecurityController) ValidateOTP(ctx *gin.Context) {
	otp := ctx.Query("otp")
	email := ctx.Query("email")
	if otp == "" || email == "" {
		slog.Error("Failed to validate OTP")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to validate OTP"})
		return
	}
	validated, err := ctrl.service.ValidateOTP(otp, email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = ctrl.userClient.CreateUserOTP(context.Background(), &pb.UserRequest{Email: email, Password: otp})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	slog.Info("OTP successfully validated")
	ctx.JSON(http.StatusOK, validated)
}
