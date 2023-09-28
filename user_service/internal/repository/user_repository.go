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
		slog.Error("failed to get user from database: %v\n", err.Error())
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
		slog.Errorf("failed to fetch user: %v\n", err.Error())
		return err
	}
	err = repo.dbClient.Save(&user).Error
	if err != nil {
		slog.Errorf("failed to update password: %v\n", err.Error())
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

	if user.Role == "Admin" {
		return nil
	}

	err = repo.dbClient.Debug().
		Model(models.User{}).
		Where("email = ?", userEmail).
		Update("role", "Admin").Error

	return err
}
