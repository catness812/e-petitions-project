package repository

import (
	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
	"github.com/catness812/e-petitions-project/petition_service/pkg/database/postgres"
	"github.com/gookit/slog"
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

func (repo *PetitionRepository) SaveVote(Vote *models.Vote) error {
	err := repo.db.Create(Vote).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PetitionRepository) UpdateStatus(id uint, statusID uint) error {
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

func (repo *PetitionRepository) Delete(id uint) error {
	err := repo.db.Unscoped().Where("id = ?", id).Delete(&models.Petition{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PetitionRepository) GetStatusByTitle(title string) (models.Status, error) {
	var status models.Status
	err := repo.db.Where("title = ?", title).First(&status).Error
	if err != nil {
		return models.Status{}, err
	}
	return status, nil
}

func (repo *PetitionRepository) GetByID(id uint) (models.Petition, error) {
	var petition models.Petition
	err := repo.db.Where("id = ?", id).Preload("Status").First(&petition).Error
	if err != nil {
		return petition, err
	}
	return petition, nil
}

func (repo *PetitionRepository) CheckIfExists(id uint) error {
	var petitions models.Petition
	if err := repo.db.Where("id = ?", id).First(&petitions).Error; err != nil {
		slog.Errorf("Couldn't find petition: %v", err.Error())
		return err
	}

	slog.Infof("petition found")
	return nil
}
