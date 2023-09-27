package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserController interface {
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

func NewUserController(service IUserService) IUserController {

	return &userController{
		service: service,
	}
}

type userController struct {
	service IUserService
}

func (c *userController) GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nice")

}

func (c *userController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nice")

}

func (c *userController) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nice")

}

func (c *userController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nice")
}
