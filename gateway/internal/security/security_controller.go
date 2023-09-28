package security

import (
	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

	ctx.Header("authorization", "Bearer "+tokens.AccessToken)
	ctx.Header("refresh-Token", tokens.RefreshToken)

	ctx.JSON(http.StatusOK, "nice login")

}

func (c *securityController) Refresh(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokens, err := c.service.Refresh(token[1])

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Header("authorization", "Bearer "+tokens.AccessToken)
	ctx.Header("refresh-Token", tokens.RefreshToken)

	ctx.JSON(http.StatusOK, "nice refresh")
}
