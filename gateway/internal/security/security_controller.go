package security

import (
	"net/http"

	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gin-gonic/gin"
)

type ISecurityController interface {
	Login(ctx *gin.Context)
	Refresh(ctx *gin.Context)
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

	ctx.Header("authorization", tokens.AccessToken)
	ctx.Header("refresh-Token", tokens.RefreshToken)

	ctx.JSON(http.StatusOK, "User logged in successfully")

}

func (c *securityController) Refresh(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")

	tokens, message, err := c.service.Refresh(authorization)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": message})
		return
	}
	ctx.Header("authorization", tokens.AccessToken)
	ctx.Header("refresh-Token", tokens.RefreshToken)

	ctx.JSON(http.StatusOK, "Refresh Token refreshed successfully")
}
