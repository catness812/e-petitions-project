package user

import (
	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"net/http"
	"strconv"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	res, err := c.service.GetByEmail(request.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	slog.Info("GetUserByEmail request successful")
	ctx.JSON(http.StatusOK, gin.H{"user": res})

}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	pid, err := strconv.ParseUint(ctx.Param("uid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the user id from param: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get the user id", "error": err})
	}

	email, err := c.service.GetByID(uint32(pid))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	slog.Info("GetUserByID request successful")
	ctx.JSON(http.StatusOK, gin.H{"email": email})

}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	res, err := c.service.Delete(request.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": res, "error": err.Error()})
		return
	}

	slog.Infof("DeleteUser request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.service.Create(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": res, "error": err.Error()})
		return
	}

	slog.Infof("CreateUser request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
func (c *UserController) OTPCreateUser(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.service.Create(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": res, "error": err.Error()})
		return
	}

	slog.Infof("OTP CreateUser request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP User created successfully"})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user model.UserCredentials
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := c.service.Update(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": res, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (c *UserController) AddAdmin(ctx *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		slog.Errorf("Invalid request format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	res, err := c.service.AddAdmin(request.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": res, "error": err.Error()})
		return
	}

	slog.Errorf("AddAdmin request successful")
	ctx.JSON(http.StatusOK, gin.H{"message": "Admin added successfully"})
}
