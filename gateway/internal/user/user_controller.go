package user

import (
	"fmt"
	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/slog"
	"log"
	"net/http"
)

type IUserService interface {
	GetByEmail(email string) (model.User, error)
	GetByID(id string) (string, error)
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

func (c *UserController) GetUserByEmail(ctx *fiber.Ctx) error {
	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		slog.Errorf("Invalid request format: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	res, err := c.service.GetByEmail(request.Email)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Info("GetUserByEmail request successful")
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *UserController) GetUserByID(ctx *fiber.Ctx) error {
	pid := ctx.Params("uid")

	email, err := c.service.GetByID(pid)
	fmt.Print(err)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Info("GetUserByID request successful")
	return ctx.JSON(email)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		slog.Errorf("Invalid request format: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := c.service.Delete(request.Email)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Infof("DeleteUser request successful")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	var user model.UserCredentials

	err := ctx.BodyParser(&user)
	if err != nil {
		log.Printf("Invalid request format: %v", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = c.service.Create(user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("CreateUser request successful")
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "User created successfully"})
}
func (c *UserController) OTPCreateUser(ctx *fiber.Ctx) error {
	var user model.UserCredentials
	if err := ctx.BodyParser(&user); err != nil {
		slog.Errorf("Invalid request format: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	_, err := c.service.Create(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Infof("OTP CreateUser request successful")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OTP User created successfully"})
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	var user model.UserCredentials
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	_, err := c.service.Update(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User updated successfully"})
}

func (c *UserController) AddAdmin(ctx *fiber.Ctx) error {
	var request struct {
		Email string `json:"email"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		slog.Errorf("Invalid request format: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	_, err := c.service.AddAdmin(request.Email)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Errorf("AddAdmin request successful")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Admin added successfully"})
}
