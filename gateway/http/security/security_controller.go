package security

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ISecurityController interface {
	Login(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}

func NewSecurityController(service ISecurityService) ISecurityController {
	log.Print("Creating new user controller")

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
