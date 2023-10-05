package petition

import (
	"context"

	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition/pb"
	"github.com/catness812/e-petitions-project/gateway/model"
)

type IPetitionRepository interface {
	CreatePetition(model.CreatePetition) (uint32, error)
	GetPetitionById(petitionID uint32) (model.Petition, error)
	GetPetitions(page uint32, limit uint32) ([]model.Petition, error)
	UpdatePetitionStatus(id uint32, status string)
	DeletePetition(petitionID uint32) error
	ValidatePetitionID(petitionID uint32) error
	CreateVote(userID uint32, petitionID uint32) error
	GetUserPetitions(userID uint32, page uint32, limit uint32) ([]model.Petition, error)
	GetUserVotedPetitions(userID uint32, page uint32, limit uint32) ([]model.Petition, error)
}

func NewPetitionRepository(c *config.Config, client pb.PetitionServiceClient) (IPetitionRepository, error) {

	us := &petitionRepository{
		cfg:    c,
		client: client,
	}

	return us, nil
}

type petitionRepository struct {
	cfg    *config.Config
	client pb.PetitionServiceClient
}

func mapPetition(pbPetition *pb.Petition) model.Petition {
	var petition model.Petition

	petition.Title = pbPetition.Title
	petition.Category = pbPetition.Category
	petition.Description = pbPetition.Description
	petition.Image = pbPetition.Image
	petition.UserID = pbPetition.UserId
	petition.Status.ID = pbPetition.Status.Id
	petition.Status.Title = pbPetition.Status.Title

	return petition
}

func mapCreatePetition(pbPetition *pb.Petition) model.CreatePetition {
	var petition model.CreatePetition

	petition.Title = pbPetition.Title
	petition.Description = pbPetition.Description
	petition.Image = pbPetition.Image
	petition.Category = pbPetition.Category
	petition.UserID = pbPetition.UserId
	petition.VoteGoal = pbPetition.VoteGoal

	return petition
}

func (repo *petitionRepository) CreatePetition(petition model.CreatePetition) (uint32, error) {
	resp, err := repo.client.CreatePetition(context.Background(), &pb.CreatePetitionRequest{
		Title:       petition.Title,
		Description: petition.Description,
		Image:       petition.Image,
		UserId:      petition.UserID,
		Category:    petition.Category,
		VoteGoal:    petition.VoteGoal,
	})

	if err != nil {
		return 0, nil
	}
	return resp.Id, nil

}

func (repo *petitionRepository) GetPetitionById(petitionID uint32) (model.Petition, error) {
	var petition model.Petition
	resp, err := repo.client.GetPetitionById(context.Background(), &pb.PetitionId{
		Id: petitionID,
	})

	if err != nil {
		return petition, err
	}

	petition = mapPetition(resp)

	return petition, nil

}

func (repo *petitionRepository) GetPetitionByID(query model.PaginationQuery) ([]model.Petition, error) {
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

// func (repo *petitionRepository) UpdatePetition(id uint32, status string) (string, error) {
// 	resp, err := repo.client.UpdatePetition(context.Background(), &pb.UpdatePetitionRequest{
// 		Id:     id,
// 		Status: status,
// 	})
// 	var message string
// 	if err != nil {
// 		return message, nil
// 	}

// 	message = resp.Message

// 	return message, nil

// }

// func (repo *petitionRepository) DeletePetition(id uint32) error {
// 	_, err := repo.client.DeletePetition(context.Background(), &pb.DeletePetitionRequest{
// 		Id: id,
// 	})

// 	return err
// }
