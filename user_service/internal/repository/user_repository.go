package repository

import (
	"fmt"

	"github.com/catness812/e-petitions-project/user_service/internal/models"
	"github.com/gookit/slog"
	"gorm.io/gorm"
)

type UserRepository struct {
	dbClient *gorm.DB
}

func NewUserRepository(dbClient *gorm.DB) *UserRepository {
	return &UserRepository{
		dbClient: dbClient,
	}
}

func (repo *UserRepository) Create(user *models.User) error {
	err := repo.dbClient.Debug().Model(models.User{}).Create(user).Error
	if err != nil {
		slog.Errorf("user failed to insert in database: %v\n", err.Error())
		return err
	}
	return nil
}

func (repo *UserRepository) Delete(userEmail string) error {

	var user models.User
	err := repo.dbClient.Debug().
		Model(models.User{}).
		Where("email = ?", userEmail).
		First(&user).Error

	if err != nil {
		slog.Error("failed to query user from database: %v\n", err.Error())
		return err
	}

	err = repo.dbClient.Debug().
		Delete(&user).Error

	if err != nil {
		slog.Error("failed to delete user from database: %v\n", err.Error())
		return err
	}

	return nil
}

func (repo *UserRepository) GetUserByEmail(userEmail string) (*models.User, error) {
	user := &models.User{}
	err := repo.dbClient.Debug().Where("email = ?", userEmail).First(user).Error
	if err != nil {
		slog.Errorf("failed to get user from database: %v\n", err.Error())
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) ValidateUserExistence(userEmail string) (*models.User, error) {
	user := &models.User{}
	err := repo.dbClient.Debug().Where("email = ?", userEmail).First(user).Error

	if err == gorm.ErrRecordNotFound {
		slog.Info("User doesn't exist: %v\n", err.Error())
		return nil, err
	} else if err != nil {
		slog.Errorf("Error fetching user: %v\n", err.Error())
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserEmailById(userID uint) (string, error) {
	var userEmail string
	err := repo.dbClient.Debug().Model(&models.User{}).Where("id = ?", userID).Pluck("email", &userEmail).Error
	if err != nil {
		slog.Errorf("failed to get user email from database %v\n", err.Error())
		return "", err
	}
	return userEmail, nil
}

func (repo *UserRepository) UpdatePasswordByEmail(user *models.User) error {
	existingUser := &models.User{}
	err := repo.dbClient.Where("email = ?", user.Email).First(&existingUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user with email %s not found", user.Email)
		}
		slog.Errorf("failed to fetch user: %v\n", err.Error())
		return err
	}
	existingUser.Password = user.Password
	err = repo.dbClient.Save(&existingUser).Error
	if err != nil {
		slog.Errorf("failed to update password: %v\n", err.Error())
		return err
	}
	return nil
}

func (repo *UserRepository) UpdateUser(user *models.User) error {
	existingUser := &models.User{}

	err := repo.dbClient.Where("id = ?", user.Id).First(&existingUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user with ID %d not found", user.Id)
		}
		slog.Errorf("failed to fetch user: %v\n", err.Error())
		return err
	}
	existingUser.Password = user.Password
	existingUser.HasAccount = user.HasAccount

	err = repo.dbClient.Save(&existingUser).Error
	if err != nil {
		slog.Errorf("failed to update user: %v\n", err.Error())
		return err
	}

	return nil
}

func (repo *UserRepository) CheckUserExistence(userEmail string) bool {
	var user models.User
	err := repo.dbClient.Debug().Model(models.User{}).Where("email = ?", userEmail).First(&user).Error
	return err == nil
}

func (repo *UserRepository) AddAdminRole(userEmail string) error {

	user := &models.User{}
	err := repo.dbClient.Debug().
		Model(models.User{}).
		Where("email = ?", userEmail).
		First(user).Error
	if err != nil {
		slog.Errorf("failed to fetch user: %v\n", err.Error())
		return err
	}

	if user.Role == "admin" {
		return nil
	}

	err = repo.dbClient.Debug().
		Model(models.User{}).
		Where("email = ?", userEmail).
		Update("role", "admin").Error

	return err
}
