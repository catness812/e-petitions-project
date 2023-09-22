package service

import (
	"github.com/catness812/e-petitions-project/internal/models"
	"github.com/catness812/e-petitions-project/internal/util"
	"gorm.io/gorm"
	"log"
)

type IPetitionRepository interface {
	Save(petition *models.Petition) error
	GetAll(pagination util.Pagination) []models.Petition
}

type PetitonService struct {
	repo IPetitionRepository
}

func InitPetitionService(repo IPetitionRepository) *PetitonService {
	return &PetitonService{
		repo: repo,
	}
}

func (svc *PetitonService) CreateNew(petition models.Petition) (uint, error) {
	if err := svc.repo.Save(&petition); err != nil {
		return 0, err
	} else {
		return petition.ID, nil
	}
}

func (svc *PetitonService) GetAll(pagination util.Pagination) []models.Petition {
	return svc.repo.GetAll(pagination)
}

func Update(db *gorm.DB, petition models.Petition) error {
	var existingPetition models.Petition
	if err := db.First(&existingPetition, petition.ID).Error; err != nil {
		log.Fatalf("Failed to find petition: %d\n", err)
		return err
	}

	existingPetition.Title = petition.Title
	existingPetition.Category = petition.Category
	existingPetition.Description = petition.Description
	existingPetition.Image = petition.Image

	if err := db.Save(&existingPetition).Error; err != nil {
		log.Fatalf("Failed to save updated petition: %d\n", err)
		return err
	}

	return nil
}

func Delete(db *gorm.DB, petition models.Petition) error {
	err := db.Delete(petition.ID).Error
	if err != nil {
		log.Fatalf("Failed to delete petition: %d\n", err)
		return err
	}
	return nil
}
