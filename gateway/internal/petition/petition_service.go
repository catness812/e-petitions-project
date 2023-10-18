package petition

import "github.com/catness812/e-petitions-project/gateway/model"

type IPetitionService interface {
	CreatePetition(model.CreatePetition) (uint32, error)
	GetPetitionByID(petitionID uint32) (model.Petition, error)
	GetPetitions(page uint32, limit uint32) ([]model.Petition, error)
	UpdatePetitionStatus(id uint32, status string) error
	DeletePetition(petitionID uint32) error
	ValidatePetitionID(petitionID uint32) error
	CreateVote(userID uint32, petitionID uint32) error
	GetUserPetitions(userID uint32, page uint32, limit uint32) ([]model.Petition, error)
	GetUserVotedPetitions(userID uint32, page uint32, limit uint32) ([]model.Petition, error)
}

func NewPetitionService(repo IPetitionRepository) (IPetitionService, error) {
	return &petitionService{
		repo: repo,
	}, nil
}

type petitionService struct {
	repo IPetitionRepository
}

func (svc *petitionService) CreatePetition(petition model.CreatePetition) (uint32, error) {
	return svc.repo.CreatePetition(petition)
}

func (svc *petitionService) GetPetitions(page uint32, limit uint32) ([]model.Petition, error) {
	return svc.repo.GetPetitions(page, limit)
}

func (svc *petitionService) GetPetitionByID(petitionID uint32) (model.Petition, error) {
	return svc.repo.GetPetitionByID(petitionID)
}

func (svc *petitionService) UpdatePetitionStatus(id uint32, status string) error {
	return svc.repo.UpdatePetitionStatus(id, status)
}

func (svc *petitionService) DeletePetition(petitionID uint32) error {
	return svc.repo.DeletePetition(petitionID)
}

func (svc *petitionService) ValidatePetitionID(petitionID uint32) error {
	return svc.repo.ValidatePetitionID(petitionID)
}

func (svc *petitionService) CreateVote(userID uint32, petitionID uint32) error {
	return svc.repo.CreateVote(userID, petitionID)
}

func (svc *petitionService) GetUserPetitions(userID uint32, page uint32, limit uint32) ([]model.Petition, error) {
	return svc.repo.GetUserPetitions(userID, page, limit)
}

func (svc *petitionService) GetUserVotedPetitions(userID uint32, page uint32, limit uint32) ([]model.Petition, error) {
	return svc.repo.GetUserVotedPetitions(userID, page, limit)
}
