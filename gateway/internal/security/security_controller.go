package security

import (
	"net/http"

	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gin-gonic/gin"
)

type ISecurityController interface {
	Login(ctx *gin.Context)
	Refresh(ctx *gin.Context)
	SendOTP(ctx *gin.Context)
}

func NewSecurityController(service ISecurityService) ISecurityController {

	return &securityController{
		service: service,
	}
}

type securityController struct {
	service ISecurityService
}

func (c *securityController) Login(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokens, err := c.service.Login(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("authorization", "Bearer "+tokens.AccessToken)
	ctx.Header("refresh-Token", tokens.RefreshToken)

	ctx.JSON(http.StatusOK, "nice login")

}

func (c *securityController) Refresh(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")

	tokens, err := c.service.Refresh(authorization)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("authorization", tokens.AccessToken)
	ctx.Header("refresh-Token", tokens.RefreshToken)

	ctx.JSON(http.StatusOK, "nice refresh")
}

func (c *securityController) SendOTP(ctx *gin.Context) {
	var email model.OTPInfo
	err := ctx.BindJSON(&email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := c.service.SendOTP(email)

	if message != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": message})
	}

	ctx.JSON(http.StatusOK, "OTP sent successfully")
}
