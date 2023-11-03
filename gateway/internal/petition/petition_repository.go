package petition

import (
	"context"
	"errors"
	"time"

	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition/pb"
	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gookit/slog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IPetitionRepository interface {
	CreatePetition(model.CreatePetition) (string, error)
	GetPetitionByID(petitionUUID string) (model.Petition, error)
	GetPetitions(page uint32, limit uint32) ([]model.Petition, error)
	UpdatePetitionStatus(uuid string, status string) error
	DeletePetition(petitionUUID string) error
	ValidatePetitionID(petitionUUID string) error
	CreateVote(userUUID string, petitionUUID string) error
	GetUserPetitions(userUUID string, page uint32, limit uint32) ([]model.Petition, error)
	GetUserVotedPetitions(userUUID string, page uint32, limit uint32) ([]model.Petition, error)
	GetAllSimilarPetitions(title string) ([]model.Petition, error)
	SearchPetitionsByTitle(title string, page uint32, limit uint32) ([]model.Petition, error)
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

	if pbPetition.Uuid == "0" {
		slog.Printf("Failed to get status value ", pbPetition.UserId)
	} else {
		petition.PetitionUUID = pbPetition.Uuid
	}
	if pbPetition.Title == "" {
		slog.Printf("Failed to get status value ", pbPetition.Title)
	} else {
		petition.Title = pbPetition.Title
	}
	if pbPetition.Category == "" {
		slog.Printf("Failed to get status value ", pbPetition.Category)
	} else {
		petition.Category = pbPetition.Category
	}
	if pbPetition.Description == "" {
		slog.Printf("Failed to get status value ", pbPetition.Description)
	} else {
		petition.Description = pbPetition.Description
	}
	petition.Image = pbPetition.Image
	if pbPetition.UserId == "0" {
		slog.Printf("Failed to get status value ", pbPetition.UserId)
	} else {
		petition.UserUUID = pbPetition.UserId
	}
	if pbPetition.AuthorName == "" {
		slog.Printf("Failed to get author name %v", pbPetition.UserId)

	} else {
		petition.AuthorName = pbPetition.AuthorName
	}
	if pbPetition.Status == nil {
		slog.Printf("Failed to get status value ", pbPetition.Status)
	} else {
		petition.Status.UUID = pbPetition.Status.Uuid
		petition.Status.Status = pbPetition.Status.Title
	}
	petition.VoteGoal = pbPetition.VoteGoal
	petition.CurrentVotes = pbPetition.CurrentVotes

	expDate := pbPetition.ExpDate.AsTime()
	petition.ExpDate = time.Time.Format(expDate, time.RFC3339) // Format as RFC3339 or your desired format

	updDate := pbPetition.UpdatedAt.AsTime()
	petition.UpdatedAt = time.Time.Format(updDate, time.RFC3339) // Format as RFC3339 or your desired format

	crtDate := pbPetition.CreatedAt.AsTime()
	petition.CreatedAt = time.Time.Format(crtDate, time.RFC3339) // Format as RFC3339 or your desired format

	return petition
}

func mapPetitionSuggestion(pbPetition *pb.PetitionInfo) model.Petition {
	var petition model.Petition
	if pbPetition.Uuid == "0" {
		slog.Printf("Failed to get status value ", pbPetition.Uuid)
	} else {
		petition.PetitionUUID = pbPetition.Uuid
	}
	if pbPetition.Title == "" {
		slog.Printf("Failed to get status value ", pbPetition.Title)
	} else {
		petition.Title = pbPetition.Title
	}
	if pbPetition.UserId == "0" {
		slog.Printf("Failed to get status value ", pbPetition.UserId)
	} else {
		petition.UserUUID = pbPetition.UserId
	}

	return petition
}

func (repo *petitionRepository) CreatePetition(petition model.CreatePetition) (string, error) {
	expDate, err := time.Parse(time.RFC3339, petition.ExpDate)
	if err != nil {
		slog.Errorf("Failed to parse expDate: %v", err)
		return "", err
	}
	expDateTimestamp := timestamppb.New(expDate)
	if expDateTimestamp.Seconds == 0 && expDateTimestamp.Nanos == 0 {
		slog.Errorf("Failed to convert time to Timestamp")
		return "", errors.New("Failed to convert time to Timestamp ")
	}
	resp, err := repo.client.CreatePetition(context.Background(), &pb.CreatePetitionRequest{
		Title:       petition.Title,
		Description: petition.Description,
		Image:       petition.Image,
		UserId:      petition.UserUUID,
		Category:    petition.Category,
		VoteGoal:    petition.VoteGoal,
		ExpDate:     expDateTimestamp,
	})

	if err != nil {
		slog.Errorf("Failed to create petition: ", err)
		return "", nil
	}
	return resp.Uuid, nil

}

func (repo *petitionRepository) GetPetitionByID(petitionID string) (model.Petition, error) {
	var petition model.Petition
	resp, err := repo.client.GetPetitionById(context.Background(), &pb.PetitionId{
		Uuid: petitionID,
	})

	if err != nil {
		slog.Infof("Failed to get petition: ", err)
		return petition, err
	}

	slog.Info("GetPetitionById: \t%v", resp)

	petition = mapPetition(resp)

	return petition, nil

}

func (repo *petitionRepository) GetPetitions(page uint32, limit uint32) ([]model.Petition, error) {
	var petitions []model.Petition
	resp, err := repo.client.GetPetitions(context.Background(), &pb.GetPetitionsRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		slog.Errorf("Failed to get all petitions: ", err)
		return petitions, err
	}

	for _, pbPetiton := range resp.Petitions {
		petitions = append(petitions, mapPetition(pbPetiton))
	}

	return petitions, nil
}

func (repo *petitionRepository) UpdatePetitionStatus(id string, status string) error {
	_, err := repo.client.UpdatePetitionStatus(context.Background(), &pb.UpdatePetitionStatusRequest{
		Uuid:   id,
		Status: status,
	})
	if err != nil {
		slog.Errorf("Failed to update petition status: ", err)
		return err
	}
	return nil
}

func (repo *petitionRepository) DeletePetition(petitionID string) error {
	_, err := repo.client.DeletePetition(context.Background(), &pb.PetitionId{
		Uuid: petitionID,
	})
	if err != nil {
		slog.Errorf("Failed to delete petition: ", err)
		return err
	}

	return nil
}

func (repo *petitionRepository) ValidatePetitionID(petitionID string) error {
	_, err := repo.client.ValidatePetitionId(context.Background(), &pb.PetitionId{
		Uuid: petitionID,
	})

	if err != nil {
		slog.Errorf("Failed to validate petition id: ", err)
		return err
	}
	return nil
}

func (repo *petitionRepository) CreateVote(userID string, petitionID string) error {
	_, err := repo.client.CreateVote(context.Background(), &pb.CreateVoteRequest{
		PetitionId: petitionID,
		UserId:     userID,
	})
	if err != nil {
		slog.Errorf("Failed to sign a petition: ", err)
		return err
	}
	return nil
}

func (repo *petitionRepository) GetUserPetitions(userID string, page uint32, limit uint32) ([]model.Petition, error) {

	resp, err := repo.client.GetUserPetitions(context.Background(), &pb.GetUserPetitionsRequest{
		UserId: userID,
		Page:   page,
		Limit:  limit,
	})
	var petitions []model.Petition

	if err != nil {
		slog.Errorf("Failed to get user created petitions: ", err)
		return petitions, err

	}

	for _, pbPetiton := range resp.Petitions {
		petitions = append(petitions, mapPetition(pbPetiton))

	}

	return petitions, nil
}

func (repo *petitionRepository) GetUserVotedPetitions(userID string, page uint32, limit uint32) ([]model.Petition, error) {
	resp, err := repo.client.GetUserVotedPetitions(context.Background(), &pb.GetUserVotedPetitionsRequest{
		UserId: userID,
		Page:   page,
		Limit:  limit,
	})
	var petitions []model.Petition

	if err != nil {
		slog.Errorf("Failed to get user voted petitions: ", err)
		return nil, err
	}

	for _, grpcPetition := range resp.Petitions {
		petition := mapPetition(grpcPetition)
		petitions = append(petitions, petition)
	}

	return petitions, nil
}

func (repo *petitionRepository) GetAllSimilarPetitions(title string) ([]model.Petition, error) {
	resp, err := repo.client.GetAllSimilarPetitions(context.Background(), &pb.PetitionSuggestionRequest{
		Title: title,
	})
	var petitions []model.Petition
	if err != nil {
		slog.Errorf("Failed to get all similar petitions: ", err)
		return nil, err
	}
	for _, grpcPetition := range resp.SuggestedPetitions {
		petition := mapPetitionSuggestion(grpcPetition)
		petitions = append(petitions, petition)
	}

	return petitions, nil
}

func (repo *petitionRepository) SearchPetitionsByTitle(title string, page uint32, limit uint32) ([]model.Petition, error) {

	resp, err := repo.client.SearchPetitionsByTitle(context.Background(), &pb.SearchPetitionsByTitRequest{
		Title: title,
		Page:  page,
		Limit: limit,
	})
	var petitions []model.Petition
	if err != nil {
		slog.Errorf("Failed to get all similar petitions: ", err)
		return nil, err
	}
	for _, grpcPetition := range resp.SuggestedPetitions {
		petition := mapPetitionSuggestion(grpcPetition)
		petitions = append(petitions, petition)
	}

	return petitions, nil

}
