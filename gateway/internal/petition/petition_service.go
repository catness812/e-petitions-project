package petition

import "github.com/catness812/e-petitions-project/gateway/model"

type IPetitionService interface {
	CreatePetition(model.CreatePetition) (string, error)
	GetPetitionByID(petitionID string) (model.Petition, error)
	GetPetitions(page uint32, limit uint32) ([]model.Petition, error)
	UpdatePetitionStatus(id string, status string) error
	DeletePetition(petitionID string) error
	ValidatePetitionID(petitionID string) error
	CreateVote(userID string, petitionID string) error
	GetUserPetitions(userID string, page uint32, limit uint32) ([]model.Petition, error)
	GetUserVotedPetitions(userID string, page uint32, limit uint32) ([]model.Petition, error)
	GetAllSimilarPetitions(title string) ([]model.Petition, error)
	SearchPetitionsByTitle(title string, page uint32, limit uint32) ([]model.Petition, error)
}

func NewPetitionService(repo IPetitionRepository) (IPetitionService, error) {
	return &petitionService{
		repo: repo,
	}, nil
}

type petitionService struct {
	repo IPetitionRepository
}

func (svc *petitionService) CreatePetition(petition model.CreatePetition) (string, error) {
	return svc.repo.CreatePetition(petition)
}

func (svc *petitionService) GetPetitions(page uint32, limit uint32) ([]model.Petition, error) {
	return svc.repo.GetPetitions(page, limit)
}

func (svc *petitionService) GetPetitionByID(petitionID string) (model.Petition, error) {
	return svc.repo.GetPetitionByID(petitionID)
}

func (svc *petitionService) UpdatePetitionStatus(id string, status string) error {
	return svc.repo.UpdatePetitionStatus(id, status)
}
func (svc *petitionService) UpdatePetition(petition model.UpdatePetition) error {
	return svc.repo.UpdatePetition(petition)
}

func (svc *petitionService) DeletePetition(petitionID string) error {
	return svc.repo.DeletePetition(petitionID)
}

func (svc *petitionService) ValidatePetitionID(petitionID string) error {
	return svc.repo.ValidatePetitionID(petitionID)
}

func (svc *petitionService) CreateVote(userID string, petitionID string) error {
	return svc.repo.CreateVote(userID, petitionID)
}

func (svc *petitionService) GetUserPetitions(userID string, page uint32, limit uint32) ([]model.Petition, error) {
	return svc.repo.GetUserPetitions(userID, page, limit)
}

func (svc *petitionService) GetUserVotedPetitions(userID string, page uint32, limit uint32) ([]model.Petition, error) {
	return svc.repo.GetUserVotedPetitions(userID, page, limit)
}

func (svc *petitionService) GetAllSimilarPetitions(title string) ([]model.Petition, error) {
	return svc.repo.GetAllSimilarPetitions(title)
}

func (svc *petitionService) SearchPetitionsByTitle(title string, page uint32, limit uint32) ([]model.Petition, error) {
	return svc.repo.SearchPetitionsByTitle(title, page, limit)
}
