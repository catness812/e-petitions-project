package petition

import "github.com/catness812/e-petitions-project/gateway/model"

type IPetitionService interface {
	CreatePetition(model.Petition) (uint32, error)
	GetPetitions(model.PaginationQuery) ([]model.Petition, error)
	UpdatePetition(id uint32, status string) (string, error)
	DeletePetition(id uint32) (string, error)
}

func NewPetitionService(repo IPetitionRepository) (IPetitionService, error) {
	return &petitionService{
		repo: repo,
	}, nil
}

type petitionService struct {
	repo IPetitionRepository
}

func (svc *petitionService) CreatePetition(petition model.Petition) (uint32, error) {
	return svc.repo.CreatePetition(petition)
}

func (svc *petitionService) GetPetitions(query model.PaginationQuery) ([]model.Petition, error) {
	return svc.repo.GetPetitions(query)
}

func (svc *petitionService) UpdatePetition(id uint32, status string) (string, error) {
	return svc.repo.UpdatePetition(id, status)
}

func (svc *petitionService) DeletePetition(id uint32) (string, error) {
	return svc.repo.DeletePetition(id)
}
