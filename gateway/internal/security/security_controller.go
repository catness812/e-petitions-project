package security

import (
	"github.com/gookit/slog"
	"net/http"

	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gin-gonic/gin"
)

type ISecurityService interface {
	Login(loginUser model.UserCredentials) (model.Tokens, error)
	Refresh(token string) (model.Tokens, error)
	SendOTP(email string) (string, error)
	ValidateOTP(otp, mail string) error
}

type SecurityController struct {
	service ISecurityService
}

func NewSecurityController(service ISecurityService) *SecurityController {
	return &SecurityController{service: service}
}

func (ctrl *SecurityController) Login(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		slog.Errorf("Could not login user: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "Could not login user"})
		return
	}
	tokens, err := ctrl.service.Login(user)
	if err != nil {
		slog.Errorf("Could not login user: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "Could not login user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "User successfully logged in",
		"access-token":  tokens.AccessToken,
		"refresh-token": tokens.RefreshToken,
	})
}

func (ctrl *SecurityController) Refresh(ctx *gin.Context) {
	type refreshToken struct {
		Token string `json:"refreshToken"`
	}
	var rt refreshToken
	err := ctx.BindJSON(&rt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Could not refresh user session",
		})
		return
	}
	tokens, err := ctrl.service.Refresh(rt.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Could not refresh user session",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":       "User successfully logged in",
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	slog.Info(email)
	_, err = ctrl.service.SendOTP(email.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "Failed to send OTP"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": false, "message": "OTP sent successfully"})
}

func (ctrl *SecurityController) ValidateOTP(ctx *gin.Context) {
	otp := ctx.Query("otp")
	email := ctx.Query("email")
	if otp == "" || email == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "Failed to validate OTP"})
		return
	}
	if err := ctrl.service.ValidateOTP(otp, email); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "Failed to validate OTP"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "OTP successfully validated",
	})
}
