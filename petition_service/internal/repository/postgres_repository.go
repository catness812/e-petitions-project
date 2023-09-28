package repository

import (
	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
	"github.com/catness812/e-petitions-project/petition_service/pkg/database/postgres"
	"gorm.io/gorm"
)

type PetitionRepository struct {
	db *gorm.DB
}

func (repo *PetitionRepository) GetAll(pagination util.Pagination) []models.Petition {
	var petitions []models.Petition
	// find paginated petitions
	repo.db.Scopes(postgres.Paginate(pagination)).Preload("Status").Find(&petitions)
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

func (repo *PetitionRepository) UpdateStatus(id uint32, statusID uint) error {
	var petition models.Petition
	// first query to see if this petition exists
	result := repo.db.Where("id = ?", id).First(&petition)
	if result.Error != nil {
		return result.Error
	}
	petition.StatusID = statusID
	repo.db.Save(&petition)
	return nil
}

func (repo *PetitionRepository) Delete(id uint32) error {
	err := repo.db.Unscoped().Where("id = ?", id).Delete(&models.Petition{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PetitionRepository) GetStatusByTitle(title string) (models.Status, error) {
	var status models.Status
	result := repo.db.Where("title = ?", title).First(&status)
	if result.Error != nil {
		return models.Status{}, result.Error
	}
	return status, nil
}
