package service

import (
	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
)

type IPetitionRepository interface {
	Save(petition *models.Petition) error
	GetAll(pagination util.Pagination) []models.Petition
	UpdateStatus(id uint, statusID uint) error
	Delete(id uint) error
	GetStatusByTitle(title string) (models.Status, error)
	GetByID(id uint) (models.Petition, error)
	GetAllUserPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error)
	SaveVote(Vote *models.Vote) error
	CheckIfExists(id uint) error
	GetAllUserVotedPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error)
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
	// save with draft status when created
	status, err := svc.repo.GetStatusByTitle(models.DRAFT)
	if err != nil {
		return 0, err
	}
	petition.Status = status
	if err := svc.repo.Save(&petition); err != nil {
		return 0, err
	} else {
		return petition.ID, nil
	}
}

func (svc *PetitonService) CreateVote(vote models.Vote) error {

	if err := svc.repo.CheckIfExists(vote.PetitionID); err != nil {
		return err
	}
	if err := svc.repo.SaveVote(&vote); err != nil {
		return err
	} else {
		return nil
	}
}

func (svc *PetitonService) GetAll(pagination util.Pagination) []models.Petition {
	return svc.repo.GetAll(pagination)
}

func (svc *PetitonService) UpdateStatus(id uint, status string) error {
	// check if status exists first
	newStatus, err := svc.repo.GetStatusByTitle(status)
	if err != nil {
		return err
	}
	if err := svc.repo.UpdateStatus(id, newStatus.ID); err != nil {
		return err
	}
	return nil
}

func (svc *PetitonService) Delete(id uint) error {
	err := svc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *PetitonService) GetByID(id uint) (models.Petition, error) {
	petition, err := svc.repo.GetByID(id)
	if err != nil {
		return petition, err
	}
	return petition, nil
}

func (svc *PetitonService) GetAllUserPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error) {
	return svc.repo.GetAllUserPetitions(userID, pagination)
}

func (svc *PetitonService) GetAllUserVotedPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error) {
	return svc.repo.GetAllUserVotedPetitions(userID, pagination)
}
