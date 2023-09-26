package service

import (
	"log"

	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
)

type IPetitionRepository interface {
	Save(petition *models.Petition) error
	GetAll(pagination util.Pagination) []models.Petition
	Update(id uint32, status string) error
	Delete(id uint32) error
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

func (svc *PetitonService)Update(id uint32, status string) error {
	if err := svc.repo.Update(id, status); err != nil {
		log.Fatalf("Failed to update petition: %d\n", err)
		return err
	}
	return nil
}

func (svc *PetitonService) Delete(id uint32) error {
    err := svc.repo.Delete(id)
    if err != nil {
        log.Fatalf("Failed to delete petition: %v\n", err)
        return err
    }
    return nil
}
