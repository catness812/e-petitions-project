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

func (repo *PetitionRepository)Update(id uint32, status string) error {
	err := repo.db.Model(&models.Petition{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PetitionRepository)Delete(id uint32) error{
	err := repo.db.Delete(&models.Petition{}, id).Error
	if err != nil{
		return err
	}
	return nil
}