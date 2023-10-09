package repository

import (
	"fmt"

	"github.com/catness812/e-petitions-project/user_service/internal/models"
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
		return err
	}

	err = repo.dbClient.Debug().
		Delete(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) GetUserByEmail(userEmail string) (*models.User, error) {
	user := &models.User{}
	err := repo.dbClient.Debug().Where("email = ?", userEmail).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) ValidateUserExistence(userEmail string) (*models.User, error) {
	user := &models.User{}
	err := repo.dbClient.Debug().Where("email = ?", userEmail).First(user).Error

	if err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserEmailById(userID uint) (string, error) {
	var userEmail string
	user := &models.User{}

	err := repo.dbClient.Debug().Where("id = ?", userID).First(user).Error
	if err != nil {
		return "", err
	}
	userEmail = user.Email
	return userEmail, nil
}

func (repo *UserRepository) UpdatePasswordByEmail(user *models.User) error {
	existingUser := &models.User{}
	err := repo.dbClient.Where("email = ?", user.Email).First(&existingUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user with email %s not found", user.Email)
		}
		return err
	}
	existingUser.Password = user.Password
	err = repo.dbClient.Save(&existingUser).Error
	if err != nil {
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
		return err
	}
	existingUser.Password = user.Password
	existingUser.HasAccount = user.HasAccount

	err = repo.dbClient.Save(&existingUser).Error
	if err != nil {
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
