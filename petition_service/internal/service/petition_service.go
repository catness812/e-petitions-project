package service

import (
	"errors"
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
	UpdateCurrVotes(petitionID uint, newCurrVotes uint) error
	HasUserVoted(userID, petitionID uint) bool
}

type IPublisherRepository interface {
	PublishMessage(email string, message string) error
}

type PetitonService struct {
	petitionRepository  IPetitionRepository
	publisherRepository IPublisherRepository
}

func NewPetitionService(petRepo IPetitionRepository, pubRepo IPublisherRepository) *PetitonService {
	return &PetitonService{
		petitionRepository:  petRepo,
		publisherRepository: pubRepo,
	}
}

func (svc *PetitonService) CreateNew(petition models.Petition) (uint, error) {
	// save with draft status when created
	status, err := svc.petitionRepository.GetStatusByTitle(models.DRAFT)
	if err != nil {
		return 0, err
	}
	petition.Status = status
	if err := svc.petitionRepository.Save(&petition); err != nil {
		return 0, err
	}

	// TODO change this after getting user's id
	err = svc.publisherRepository.PublishMessage("test@email.com", "Your petition has been created")
	if err != nil {
		return 0, err
	}

	return petition.ID, nil
}

func (svc *PetitonService) CreateVote(vote models.Vote) error {
	if res := svc.petitionRepository.HasUserVoted(vote.UserID, vote.PetitionID); res {
		return errors.New("user has already voted")
	} else {
		if err := svc.petitionRepository.CheckIfExists(vote.PetitionID); err != nil {
			return err
		}
		if err := svc.petitionRepository.SaveVote(&vote); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func (svc *PetitonService) GetAll(pagination util.Pagination) []models.Petition {
	return svc.petitionRepository.GetAll(pagination)
}

func (svc *PetitonService) UpdateStatus(id uint, status string) error {
	// check if status exists first
	newStatus, err := svc.petitionRepository.GetStatusByTitle(status)
	if err != nil {
		return err
	}
	if err := svc.petitionRepository.UpdateStatus(id, newStatus.ID); err != nil {
		return err
	}
	return nil
}

func (svc *PetitonService) Delete(id uint) error {
	err := svc.petitionRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *PetitonService) GetByID(id uint) (models.Petition, error) {
	petition, err := svc.petitionRepository.GetByID(id)
	if err != nil {
		return petition, err
	}
	return petition, nil
}

func (svc *PetitonService) GetAllUserPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error) {
	return svc.petitionRepository.GetAllUserPetitions(userID, pagination)
}

func (svc *PetitonService) GetAllUserVotedPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error) {
	return svc.petitionRepository.GetAllUserVotedPetitions(userID, pagination)
}

func (svc *PetitonService) UpdateCurrVotes(petitionID uint, newCurrVotes uint) error {
	return svc.petitionRepository.UpdateCurrVotes(petitionID, newCurrVotes)
}
