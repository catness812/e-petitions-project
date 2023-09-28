package security_repository

import (
	"errors"
	"log"

	models "github.com/catness812/e-petitions-project/security_service/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DBClient *gorm.DB
}

func NewUserRepository(dbClient *gorm.DB) *UserRepository {
	return &UserRepository{
		DBClient: dbClient,
	}
}

func (repo *UserRepository) CheckIfEmailExists(mail string) bool {
	var user models.UserModel

	err := repo.DBClient.Debug().Model(models.UserModel{}).Find(&user).Where("email= ?", mail).Error

	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (repo *UserRepository) GetUserByEmail(email string) (models.UserModel, error) {
	var user models.UserModel

	err := repo.DBClient.Debug().Model(models.UserModel{}).Where("email = ?", email).First(&user).Error

	if err != nil {
		log.Printf("Invalid credentials: %v", err)
		return models.UserModel{}, err
	}
	return user, nil
}
