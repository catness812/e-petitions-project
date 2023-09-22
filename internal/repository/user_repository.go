package repository

import (
	"fmt"
	"user_service/internal/models"

	"golang.org/x/exp/slog"
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
		slog.Error("user failed to insert in database: %v\n", err)
		return err
	}
	return nil
}

func (repo *UserRepository) Delete(userEmail string) error {
	err := repo.dbClient.Debug().
		Model(models.User{}).
		Where("email = ?", userEmail).
		Delete(&models.User{}).Error
	if err != nil {
		slog.Error("failed to delete user from database: %v\n", err)
		return err
	}

	return nil
}

func (repo *UserRepository) GetUserByEmail(userEmail string) (*models.User, error) {
	user := &models.User{}
	err := repo.dbClient.Debug().Where("email = ?", userEmail).First(user).Error
	if err != nil {
		slog.Error("failed to get user from database: %v\n", err)
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) UpdatePasswordByEmail(user *models.User) error {
	err := repo.dbClient.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user with email %s not found", user.Email)
		}
		slog.Error("failed to fetch user: %v\n", err)
		return err
	}
	err = repo.dbClient.Save(&user).Error
	if err != nil {
		slog.Error("failed to update password: %v\n", err)
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
	// Fetch the user by email
	user := &models.User{}
	err := repo.dbClient.Debug().
		Model(models.User{}).
		Where("email = ?", userEmail).
		First(user).Error
	if err != nil {
		return err
	}

	if user.Role == "Admin" {
		return nil
	}

	// Update the role
	err = repo.dbClient.Debug().
		Model(models.User{}).
		Where("email = ?", userEmail).
		Update("role", "Admin").Error

	return err
}
