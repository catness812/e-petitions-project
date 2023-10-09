package service

import (
	"fmt"
	"time"

	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
)

type IPetitionRepository interface {
	Save(petition *models.Petition) error
	GetAll(pagination util.Pagination) []models.Petition
	GetAllActive() []models.Petition
	UpdateStatus(id uint, statusID uint) error
	Delete(id uint) error
	GetStatusByTitle(title string) (models.Status, error)
	GetByID(id uint) (models.Petition, error)
	GetAllUserPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error)
	SaveVote(Vote *models.Vote) error
	CheckIfExists(id uint) error
	GetAllUserVotedPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error)
	UpdateCurrVotes(petition models.Petition) error
	HasUserVoted(userID, petitionID uint) error
}

type IPublisherRepository interface {
	PublishMessage(email string, message string) error
}

type IUserRepository interface {
	GetEmailById(id uint) (string, error)
}

type PetitonService struct {
	petitionRepository  IPetitionRepository
	publisherRepository IPublisherRepository
	userRepository      IUserRepository
}

func NewPetitionService(
	petRepo IPetitionRepository,
	pubRepo IPublisherRepository,
	userRepo IUserRepository,
) *PetitonService {
	return &PetitonService{
		petitionRepository:  petRepo,
		publisherRepository: pubRepo,
		userRepository:      userRepo,
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

	// get user's email from User Service
	email, err := svc.userRepository.GetEmailById(petition.UserID)
	if err != nil {
		return 0, err
	}
	err = svc.publisherRepository.PublishMessage(email, fmt.Sprintf(`Petition "%s" has been successfully created!`, petition.Title))
	if err != nil {
		return 0, err
	}

	return petition.ID, nil
}

func (svc *PetitonService) CreateVote(vote models.Vote) error {
	// see if petition exists
	petition, err := svc.petitionRepository.GetByID(vote.PetitionID)
	if err != nil {
		return err
	}

	if err := svc.petitionRepository.HasUserVoted(vote.UserID, vote.PetitionID); err != nil {
		return err
	}

	petition.CurrVotes++
	if err := svc.petitionRepository.UpdateCurrVotes(petition); err != nil {
		return err
	}

	if err := svc.petitionRepository.SaveVote(&vote); err != nil {
		return err
	}

	// get user's email from User Service
	email, err := svc.userRepository.GetEmailById(petition.UserID)
	if err != nil {
		return err
	}

	// if the vote & petition were saved successfully, send email on vote goal
	if petition.VoteGoal == petition.CurrVotes {
		err = svc.publisherRepository.PublishMessage(email, fmt.Sprintf(
			`Petition "%s" has been reached its goal of %d votes! Congrats!`,
			petition.Title, petition.VoteGoal))
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *PetitonService) GetAll(pagination util.Pagination) []models.Petition {
	return svc.petitionRepository.GetAll(pagination)
}

func (svc *PetitonService) GetAllActive() []models.Petition {
	return svc.petitionRepository.GetAllActive()
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

func (svc *PetitonService) CheckPetitionExpiration(petition models.Petition) (string, error) {
	if time.Now().After(petition.ExpDate) {
		email, err := svc.userRepository.GetEmailById(petition.UserID)
		if err != nil {
			return "", err
		}

		err = svc.UpdateStatus(petition.ID, "ARCHIVE")
		if err != nil {
			return "", err
		}

		err = svc.publisherRepository.PublishMessage(email, fmt.Sprintf(`Petition "%s" has expired! It's been moved to your archived petitions.`, petition.Title))
		if err != nil {
			return "", err
		}
	}
	return "", nil
}
