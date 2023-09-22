package service

import (
	"errors"
	"user_service/internal/models"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
)

type IUserRepository interface {
	Create(user *models.User) error
	UpdatePasswordByEmail(user *models.User) error
	CheckUserExistence(userEmail string) bool
	Delete(userEmail string) error
	GetUserByEmail(userEmail string) (*models.User, error)
	AddAdminRole(userEmail string) error
}

type UserService struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (svc *UserService) Create(user *models.User) error {
	user.Role = "user"
	if svc.userRepo.CheckUserExistence(user.Email) {
		return errors.New("user exists")
	}

	hashedPassword, err := svc.generatePasswordHash(user.Password)
	if err != nil {
		return errors.New("can't register")
	}
	user.Password = hashedPassword
	return svc.userRepo.Create(user)
}

func (svc *UserService) generatePasswordHash(password string) (string, error) {
	const salt = 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		slog.Error("ERR: %v\n", err)
		return "", err
	}
	return string(hashedPassword), err
}

func (svc *UserService) UpdatePasswordByEmail(user *models.User) error {
	hashedPassword, err := svc.generatePasswordHash(user.Password)
	if err != nil {
		return errors.New("error generating password hash")
	}
	user.Password = hashedPassword
	err = svc.userRepo.UpdatePasswordByEmail(user)
	if err != nil {
		return errors.New("error updating password")
	}
	return nil
}

func (svc *UserService) Delete(userEmail string) error {
	err := svc.userRepo.Delete(userEmail)
	if err != nil {
		slog.Error("failed to delete user from database: %v\n", err)
		return err
	}
	return nil
}

func (svc *UserService) AddAdmin(userEmail string) error {
	err := svc.userRepo.AddAdminRole(userEmail)
	if err != nil {
		slog.Error("failed to add admin role: %v\n", err)
		return err
	}
	return nil
}

func (svc *UserService) GetUserByEmail(userEmail string) (*models.User, error) {
	user, err := svc.userRepo.GetUserByEmail(userEmail)
	if err != nil {
		slog.Error("failed to get user from database: %v\n", err)
		return nil, err
	}
	return user, nil
}