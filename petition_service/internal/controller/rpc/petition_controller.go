package rpctransport

import (
	"context"
	"errors"
	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/pb"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type IPetitionSvc interface {
	CreateNew(petition models.Petition) (uint, error)
	GetAll(pagination util.Pagination) []models.Petition
	UpdateStatus(id uint32, status string) error
	Delete(id uint32) error
}

type Server struct {
	pb.PetitionServiceServer
	PetitionService IPetitionSvc
}

func (s *Server) CreatePetition(_ context.Context, req *pb.CreatePetitionRequest) (*pb.PetitionId, error) {

	newPetition := models.Petition{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		UserID:      uint(req.UserId),
		Category:    req.Category,
		VoteGoal:    uint(req.VoteGoal),
	}

	savedPetitionID, err := s.PetitionService.CreateNew(newPetition)
	if err != nil {
		return nil, err
	}

	return &pb.PetitionId{
		Id: uint32(savedPetitionID),
	}, nil
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
			UserId: uint32(p.UserID),
		}
	}
	return &pb.GetPetitionsResponse{
		Petitions: getPetitionsResponse,
	}, nil
}

func (s *Server) UpdatePetitionStatus(_ context.Context, req *pb.UpdatePetitionStatusRequest) (*empty.Empty, error) {
	id := req.Id
	statusTitle := req.Status

	err := s.PetitionService.UpdateStatus(id, statusTitle)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	response := &empty.Empty{}
	return response, nil
}

func (s *Server) DeletePetition(_ context.Context, req *pb.PetitionId) (*empty.Empty, error) {
	id := req.Id
	err := s.PetitionService.Delete(id)
	if err != nil {
		return nil, err
	}
	response := &empty.Empty{}
	return response, nil
}
