package security

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	ctx.JSON(http.StatusOK, "nice")

}

func (c *securityController) Refresh(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nice")

}
