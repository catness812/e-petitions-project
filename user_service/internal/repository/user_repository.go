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
	err := repo.dbClient.Debug().Where("email = ?", userEmail).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) ValidateUserExistence(userEmail string) (*models.User, error) {
	user := &models.User{}
	err := repo.dbClient.Debug().Where("email = ?", userEmail).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == gorm.ErrRecordNotFound || user.UUID == "0" {
		return nil, nil
	}
	return user, nil
}

func (repo *UserRepository) GetUserEmailById(userID string) (string, error) {
	var userEmail string
	user := &models.User{}

	err := repo.dbClient.Debug().Where("uuid = ?", userID).First(&user).Error
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

	err := repo.dbClient.Where("uuid = ?", user.UUID).First(&existingUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user with ID %d not found", user.UUID)
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

func (repo *UserRepository) CheckUserExistence(userid string) (bool, error) {
	var user models.User
	err := repo.dbClient.Debug().Model(models.User{}).Where("uuid = ?", userid).First(&user).Error
	if user.UUID == "0" || err != nil {
		return false, err
	}
	return true, err
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

func (repo *UserRepository) GetAdminEmails() ([]string, error) {
	var emails []string
	err := repo.dbClient.Debug().
		Model(models.User{}).
		Where("role = ?", "admin").
		Pluck("email", &emails).
		Error

	return emails, err
}
