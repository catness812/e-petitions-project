package user

import (
	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"net/http"
	"strconv"
	"strings"
)

type IUserService interface {
	GetByEmail(email string) (model.User, error)
	GetByID(id uint32) (string, error)
	Delete(email string) (string, error)
	Create(createUser model.UserCredentials) (string, error)
	OTPCreate(createUser model.UserCredentials) (string, error)
	Update(createUser model.UserCredentials) (string, error)
	AddAdmin(email string) (string, error)
}

type UserController struct {
	service IUserService
}

func NewUserController(service IUserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetUserByEmail(ctx *gin.Context) {
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
		ctx.JSON(http.StatusNotFound, gin.H{"message": strings.Split(err.Error(), "=")[2], "error": true})
		return
	}

	slog.Info("GetUserByEmail request successful")
	ctx.JSON(http.StatusOK, gin.H{"user": res, "error": false, "message": "User found"})

}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	pid, err := strconv.ParseUint(ctx.Param("uid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the user id from param: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get the user id", "error": err})
	}

	email, err := c.service.GetByID(uint32(pid))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": strings.Split(err.Error(), "=")[2], "error": true})
		return
	}

	slog.Info("GetUserByID request successful")
	ctx.JSON(http.StatusOK, gin.H{"email ": email, "error": false, "message": "User found"})

}

func (c *UserController) DeleteUser(ctx *gin.Context) {
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
		ctx.JSON(http.StatusNotFound, gin.H{"message": strings.Split(err.Error(), "=")[2], "error": true})
		return
	}

	slog.Infof("DeleteUser request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "error": false})

}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}
	_, err = c.service.Create(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": strings.Split(err.Error(), "=")[2], "error": true})
		return
	}

	slog.Infof("CreateUser request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "error": false})
}
func (c *UserController) OTPCreateUser(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}
	_, err = c.service.Create(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": strings.Split(err.Error(), "=")[2], "error": true})
		return
	}

	slog.Infof("OTP CreateUser request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP User created successfully", "error": false})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}
	_, err = c.service.Update(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": strings.Split(err.Error(), "=")[2], "error": true})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "error": false})
}

func (c *UserController) AddAdmin(ctx *gin.Context) {
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
		ctx.JSON(http.StatusNotFound, gin.H{"message": strings.Split(err.Error(), "=")[2], "error": true})
		return
	}

	slog.Errorf("AddAdmin request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "Admin added successfully", "error": false})
}
