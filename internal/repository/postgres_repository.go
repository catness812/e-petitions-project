package repository

import (
	"github.com/catness812/e-petitions-project/internal/models"
	"github.com/catness812/e-petitions-project/internal/util"
	"github.com/catness812/e-petitions-project/pkg/database/postgres"
	"gorm.io/gorm"
)

type PetitionRepository struct {
	db *gorm.DB
}

func (repo *PetitionRepository) GetAll(pagination util.Pagination) []models.Petition {
	var petitions []models.Petition
	// find paginated petitions
	repo.db.Scopes(postgres.Paginate(pagination)).Find(&petitions)
	return petitions
}

func InitPetitionRepository(db *gorm.DB) *PetitionRepository {
	return &PetitionRepository{
		db: db,
	}
}

func (repo *PetitionRepository) Save(petition *models.Petition) error {
	err := repo.db.Create(petition).Error
	if err != nil {
		return err
	}
	return nil
}

// func Update(petition *models.Petition) error {
// 	err := postgres.Database.Model(&models.Petition{}).Where("id = ?", petition.ID).Update("status", petition.Status).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
