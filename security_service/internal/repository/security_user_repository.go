package security_repository

import (
	"errors"
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
	var user models.User

	err := repo.DBClient.Debug().Model(models.User{}).Find(&user).Where("email= ?", mail).Error

	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (repo *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := repo.DBClient.Debug().Model(models.User{}).Where("email = ?", email).First(&user).Error

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
