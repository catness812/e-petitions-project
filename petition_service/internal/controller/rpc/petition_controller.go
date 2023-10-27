package rpc

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/pb"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type IPetitionService interface {
	CreateNew(petition models.Petition) (uint, error)
	GetAll(pagination util.Pagination) []models.Petition
	UpdateStatus(id uint, status string) error
	Delete(id uint) error
	GetByID(id uint) (models.Petition, error)
	CreateVote(vote models.Vote) error
	GetAllUserPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error)
	GetAllUserVotedPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error)
	CheckPetitionExpiration(petition models.Petition) (string, error)
	GetAllSimilarPetitions(title string) ([]models.PetitionInfo, error)
	SearchPetitionsByTitle(searchTerm string, pagination util.Pagination) ([]models.PetitionInfo, error)
	ScheduleDailyCheck()
}

type Server struct {
	pb.PetitionServiceServer
	PetitionService IPetitionService
}

func (s *Server) GetPetitionById(_ context.Context, req *pb.PetitionId) (*pb.Petition, error) {
	petition, err := s.PetitionService.GetByID(uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "petition not found")
		}
		return nil, err
	}

	return &pb.Petition{
		Id:          uint32(petition.ID),
		Title:       petition.Title,
		Category:    petition.Category,
		Description: petition.Description,
		Image:       petition.Image,
		Status: &pb.Status{
			Id:    uint32(petition.Status.ID),
			Title: petition.Status.Title,
		},
		UserId:       uint32(petition.UserID),
		VoteGoal:     uint32(petition.VoteGoal),
		CreatedAt:    timestamppb.New(petition.CreatedAt),
		UpdatedAt:    timestamppb.New(petition.UpdatedAt),
		CurrentVotes: uint32(petition.CurrVotes),
		ExpDate:      timestamppb.New(petition.ExpDate),
	}, nil
}

func (s *Server) ValidatePetitionId(_ context.Context, req *pb.PetitionId) (*empty.Empty, error) {
	_, err := s.PetitionService.GetByID(uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "petition not found")
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *Server) CreatePetition(_ context.Context, req *pb.CreatePetitionRequest) (*pb.PetitionId, error) {
	newPetition := models.Petition{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		UserID:      uint(req.UserId),
		Category:    req.Category,
		VoteGoal:    uint(req.VoteGoal),
		ExpDate:     req.ExpDate.AsTime(),
	}

	savedPetitionID, err := s.PetitionService.CreateNew(newPetition)
	if err != nil {
		return nil, err
	}

	return &pb.PetitionId{
		Id: uint32(savedPetitionID),
	}, nil
}

func (s *Server) CreateVote(_ context.Context, req *pb.CreateVoteRequest) (*empty.Empty, error) {
	newVote := models.Vote{
		PetitionID: uint(req.PetitionId),
		UserID:     uint(req.UserId),
	}

	err := s.PetitionService.CreateVote(newVote)
	if err != nil {
		return nil, err
	}

	response := &empty.Empty{}
	return response, nil
}

func (s *Server) GetPetitions(_ context.Context, req *pb.GetPetitionsRequest) (*pb.GetPetitionsResponse, error) {
	pag := util.Pagination{
		Page:  int(req.Page),
		Limit: int(req.Limit),
	}

	petitions := s.PetitionService.GetAll(pag)

	getPetitionsResponse := make([]*pb.Petition, len(petitions))

	for i := range getPetitionsResponse {
		p := petitions[i]
		getPetitionsResponse[i] = &pb.Petition{
			Id:          uint32(p.ID),
			Title:       p.Title,
			Category:    p.Category,
			Description: p.Description,
			Image:       p.Image,

			Status: &pb.Status{
				Id:    uint32(p.Status.ID),
				Title: p.Status.Title,
			},
			UserId:       uint32(p.UserID),
			VoteGoal:     uint32(p.VoteGoal),
			CreatedAt:    timestamppb.New(p.CreatedAt),
			UpdatedAt:    timestamppb.New(p.UpdatedAt),
			CurrentVotes: uint32(p.CurrVotes),
			ExpDate:      timestamppb.New(p.ExpDate),
		}
	}

	return &pb.GetPetitionsResponse{
		Petitions: getPetitionsResponse,
	}, nil
}

func (s *Server) UpdatePetitionStatus(_ context.Context, req *pb.UpdatePetitionStatusRequest) (*empty.Empty, error) {
	err := s.PetitionService.UpdateStatus(uint(req.Id), req.Status)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "petition not found")
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *Server) DeletePetition(_ context.Context, req *pb.PetitionId) (*empty.Empty, error) {
	err := s.PetitionService.Delete(uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "petition not found")
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *Server) GetUserPetitions(_ context.Context, req *pb.GetUserPetitionsRequest) (*pb.GetUserPetitionsResponse, error) {
	userID := req.UserId
	pag := util.Pagination{
		Page:  int(req.Page),
		Limit: int(req.Limit),
	}

	petitions, err := s.PetitionService.GetAllUserPetitions(uint(userID), pag)
	if err != nil {
		return nil, err
	}
	getUserPetitionsResponse := make([]*pb.Petition, len(petitions))

	for i := range getUserPetitionsResponse {
		p := petitions[i]
		pStatus := &pb.Status{
			Id:    uint32(p.StatusID),
			Title: p.Status.Title,
		}
		getUserPetitionsResponse[i] = &pb.Petition{
			Id:           uint32(p.ID),
			Title:        p.Title,
			Category:     p.Category,
			Description:  p.Description,
			VoteGoal:     uint32(p.VoteGoal),
			CreatedAt:    timestamppb.New(p.CreatedAt),
			UpdatedAt:    timestamppb.New(p.UpdatedAt),
			CurrentVotes: uint32(p.CurrVotes),
			ExpDate:      timestamppb.New(p.ExpDate),
			UserId:       uint32(p.UserID),
			Status:       pStatus,
		}
	}

	return &pb.GetUserPetitionsResponse{
		Petitions: getUserPetitionsResponse,
	}, nil
}

func (s *Server) GetUserVotedPetitions(_ context.Context, req *pb.GetUserVotedPetitionsRequest) (*pb.GetUserVotedPetitionsResponse, error) {
	userID := req.UserId
	pag := util.Pagination{
		Page:  int(req.Page),
		Limit: int(req.Limit),
	}
	petitions, err := s.PetitionService.GetAllUserVotedPetitions(uint(userID), pag)
	if err != nil {
		return nil, err
	}
	getUserPetitionsResponse := make([]*pb.Petition, len(petitions))

	for i := range getUserPetitionsResponse {
		p := petitions[i]
		getUserPetitionsResponse[i] = &pb.Petition{
			Id:           uint32(p.ID),
			Title:        p.Title,
			Category:     p.Category,
			Description:  p.Description,
			VoteGoal:     uint32(p.VoteGoal),
			CreatedAt:    timestamppb.New(p.CreatedAt),
			UpdatedAt:    timestamppb.New(p.UpdatedAt),
			CurrentVotes: uint32(p.CurrVotes),
			ExpDate:      timestamppb.New(p.ExpDate),
			UserId:       uint32(p.UserID),
			Status:       &pb.Status{Id: uint32(p.Status.ID), Title: p.Status.Title},
		}
	}

	return &pb.GetUserVotedPetitionsResponse{
		Petitions: getUserPetitionsResponse,
	}, nil
}

func (s *Server) CheckIfPetitionsExpired(_ context.Context, req *pb.Petition) (*empty.Empty, error) {
	petition, err := s.PetitionService.GetByID(uint(req.Id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "petition not found")
		}
		return nil, err
	}

	if _, err := s.PetitionService.CheckPetitionExpiration(petition); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *Server) ScheduleDailyCheck() {
	s.PetitionService.ScheduleDailyCheck()
}

func (s *Server) GetAllSimilarPetitions(_ context.Context, req *pb.PetitionSuggestionRequest) (*pb.PetitionSuggestionResponse, error) {
	petitions, err := s.PetitionService.GetAllSimilarPetitions(req.Title)

	if err != nil {
		return nil, err
	}

	getAllSimilarPetitionsResponse := make([]*pb.PetitionInfo, len(petitions))

	for i := range getAllSimilarPetitionsResponse {
		p := petitions[i]
		getAllSimilarPetitionsResponse[i] = &pb.PetitionInfo{
			Id:     uint32(p.ID),
			Title:  p.Title,
			UserId: uint32(p.UserID),
		}
	}
	return &pb.PetitionSuggestionResponse{
		SuggestedPetitions: getAllSimilarPetitionsResponse,
	}, nil

}

func (s *Server) SearchPetitionsByTitle(_ context.Context, req *pb.SearchPetitionsByTitRequest) (*pb.PetitionSuggestionResponse, error) {
	pag := util.Pagination{
		Page:  int(req.Page),
		Limit: int(req.Limit),
	}
	petitions, err := s.PetitionService.SearchPetitionsByTitle(req.Title, pag)
	if err != nil {
		return nil, err
	}

	SearchPetitionsByTitleResponse := make([]*pb.PetitionInfo, len(petitions))

	for i := range SearchPetitionsByTitleResponse {
		p := petitions[i]
		SearchPetitionsByTitleResponse[i] = &pb.PetitionInfo{
			Id:     uint32(p.ID),
			Title:  p.Title,
			UserId: uint32(p.UserID),
		}
	}
	return &pb.PetitionSuggestionResponse{
		SuggestedPetitions: SearchPetitionsByTitleResponse,
	}, nil
}
