package service

import (
	"errors"
	"regexp"

	"github.com/catness812/e-petitions-project/user_service/internal/models"
	"github.com/gookit/slog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *models.User) error
	UpdatePasswordByEmail(user *models.User) error
	CheckUserExistence(userEmail string) bool
	Delete(userEmail string) error
	GetUserByEmail(userEmail string) (*models.User, error)
	AddAdminRole(userEmail string) error
	GetUserEmailById(userID uint) (string, error)
	ValidateUserExistence(userEmail string) (*models.User, error)
	UpdateUser(user *models.User) error
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
	existingUser, err := svc.userRepo.ValidateUserExistence(user.Email)
	if err != nil && existingUser == nil {
		slog.Errorf("ERR validating user existence: %v\n", err)
		return err
	} else if existingUser == nil && err == nil {
		// if user doesn't exists it creates it
		user.Role = "user"
		valid := validMailAddress(user.Email)
		if valid == false {
			slog.Info("invalid email")
			return errors.New("invalid email")
		}
		hashedPassword, err := svc.generatePasswordHash(user.Password)
		if err != nil {
			slog.Errorf("can't register: %v\n", err.Error())
			return errors.New("can't register")
		}
		user.Password = hashedPassword
		err = svc.userRepo.Create(user)
		if err != nil {
			slog.Errorf("user failed to insert in database: %v\n", err.Error())
			return err
		}
		slog.Info("User added successfully")
		return nil
	} else if err == nil && existingUser != nil {
		if !existingUser.HasAccount {
			// if user exists but was not previously registered
			existingUser.Password = user.Password
			existingUser.HasAccount = user.HasAccount
			hashedPassword, err := svc.generatePasswordHash(existingUser.Password)
			if err != nil {
				slog.Errorf("Generating paswword hash: %v\n", err.Error())
				return errors.New("error generating password hash")
			}
			existingUser.Password = hashedPassword
			slog.Info("hashed pass ", hashedPassword)
			err = svc.userRepo.UpdateUser(existingUser)
			if err != nil {
				slog.Errorf("error updating password: %v\n", err.Error())
				return errors.New("error updating password")
			}
			slog.Info("User register successfully")
			return nil
		} else {
			slog.Info("User already Exists")
			return errors.New("User already Exists")
		}
	}
	return err
}

func (svc *UserService) generatePasswordHash(password string) (string, error) {
	const salt = 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		slog.Errorf("ERR: %v\n", err)
		return "", err
	}
	return string(hashedPassword), err
}

func (svc *UserService) UpdatePasswordByEmail(user *models.User) error {
	hashedPassword, err := svc.generatePasswordHash(user.Password)
	if err != nil {
		slog.Errorf("error generating password hash: %v\n", err.Error())
		return errors.New("error generating password hash")
	}
	user.Password = hashedPassword
	slog.Info("hashed pass ", user.Password)
	err = svc.userRepo.UpdatePasswordByEmail(user)
	if err != nil {
		slog.Errorf("error updating password: %v\n", err.Error())
		return errors.New("error updating password")
	}
	return nil
}

func (svc *UserService) Delete(userEmail string) error {
	err := svc.userRepo.Delete(userEmail)
	if err != nil {
		slog.Errorf("failed to delete user from database: %v", err.Error())
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

func (svc *UserService) GetUserEmailById(userID uint) (string, error) {
	userEmail, err := svc.userRepo.GetUserEmailById(userID)
	if err == gorm.ErrRecordNotFound {
		slog.Infof("User with ID %d not found", userID)
		return "", gorm.ErrRecordNotFound
	} else if err != nil {
		slog.Errorf("Failed to fetch user from database: %v", err.Error())
		return "", err
	}
	return userEmail, nil
}

func validMailAddress(address string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	valid, err := regexp.MatchString(regex, address)
	if err != nil {
		slog.Errorf("error checking mail: %v\n", err.Error())
		return false
	}
	return valid
}
