package user

import (
	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"net/http"
	"strconv"
)

type IUserController interface {
	GetUserByEmail(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	OTPCreateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	AddAdmin(ctx *gin.Context)
}

func NewUserController(service IUserService) IUserController {

	return &userController{
		service: service,
	}
}

type userController struct {
	service IUserService
}

func (c *userController) GetUserByEmail(ctx *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request format", "error": true})
		return
	}

	res, err := c.service.GetByEmail(request.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Info("GetUserByEmail request successful")
	ctx.JSON(http.StatusOK, gin.H{"user": res, "error": false, "message": "User found"})

}

func (c *userController) GetUserByID(ctx *gin.Context) {
	pid, err := strconv.ParseUint(ctx.Param("uid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the user id from param: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get the user id", "error": err})
	}

	email, err := c.service.GetByID(uint32(pid))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Info("GetUserByID request successful")
	ctx.JSON(http.StatusOK, gin.H{"email ": email, "error": false, "message": "User found"})

}

func (c *userController) DeleteUser(ctx *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request format", "error": true})
		return
	}

	_, err := c.service.Delete(request.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Infof("DeleteUser request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "error": false})

}

func (c *userController) CreateUser(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}
	_, err = c.service.Create(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Infof("CreateUser request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "error": false})
}
func (c *userController) OTPCreateUser(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}
	_, err = c.service.Create(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Infof("OTP CreateUser request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP User created successfully", "error": false})
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}
	_, err = c.service.Update(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "error": false})
}

func (c *userController) AddAdmin(ctx *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request format", "error": true})
		return
	}

	_, err := c.service.AddAdmin(request.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Errorf("AddAdmin request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "Admin added successfully", "error": false})
}
