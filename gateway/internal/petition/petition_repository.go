package petition

import (
	"context"

	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition/pb"
	"github.com/catness812/e-petitions-project/gateway/model"
)

type IPetitionRepository interface {
	CreatePetition(model.Petition) (uint32, error)
	GetPetitions(model.PaginationQuery) ([]model.Petition, error)
	UpdatePetition(id uint32, status string) (string, error)
	DeletePetition(id uint32) (string, error)
}

func NewPetitionRepository(c config.Config, client pb.PetitionServiceClient) (IPetitionRepository, error) {

	us := &petitionRepository{
		cfg:    c,
		client: client,
	}

	return us, nil
}

type petitionRepository struct {
	cfg    config.Config
	client pb.PetitionServiceClient
}

func mapPetition(pbPetition *pb.Petition) model.Petition {
	var petition model.Petition

	petition.Title = pbPetition.Title
	petition.Category = pbPetition.Category
	petition.Description = pbPetition.Description
	petition.Image = pbPetition.Image
	petition.Status = uint32(pbPetition.Status)
	petition.UserID = uint(pbPetition.UserId)

	return petition
}

func (repo *petitionRepository) CreatePetition(petition model.Petition) (uint32, error) {
	resp, err := repo.client.CreatePetition(context.Background(), &pb.CreatePetitionRequest{
		Petition: &pb.Petition{
			PetitionId: uint32(petition.PetitionId),
			Title:      petition.Title,
			Category:   petition.Category,
			Image:      petition.Image,
			Status:     uint32(petition.Status),
			UserId:     uint32(petition.UserID),
		},
	})

	if err != nil {
		return 0, err
	}
	return resp.PetitionId, nil
}

func (repo *petitionRepository) GetPetitions(query model.PaginationQuery) ([]model.Petition, error) {
	resp, err := repo.client.GetPetitions(context.Background(), &pb.GetPetitionsRequest{
		Page:  query.Page,
		Limit: query.Limit,
	})

	var petitions []model.Petition

	if err != nil {
		return petitions, err
	}

	for _, pbpetition := range resp.Petitions {
		petitions = append(petitions, mapPetition(pbpetition))
	}

	return petitions, nil
}

func (repo *petitionRepository) UpdatePetition(id uint32, status string) (string, error) {
	resp, err := repo.client.UpdatePetition(context.Background(), &pb.UpdatePetitionRequest{
		Id:     id,
		Status: status,
	})
	var message string
	if err != nil {
		return message, nil
	}

	message = resp.Message

	return message, nil

}

func (repo *petitionRepository) DeletePetition(id uint32) (string, error) {
	resp, err := repo.client.DeletePetition(context.Background(), &pb.DeletePetitionRequest{
		Id: id,
	})

	var message string
	if err != nil {
		return message, err
	}

	message = resp.Message

	return message, nil
}
