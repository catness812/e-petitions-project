package rpctransport

import (
	"context"

	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/pb"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
)

type IPetitionSvc interface {
	CreateNew(petition models.Petition) (uint, error)
	GetAll(pagination util.Pagination) []models.Petition
	Update(id uint32, status string) error
	Delete(id uint32) error
}

type Server struct {
	pb.PetitionServiceServer
	Svc IPetitionSvc
}

func (s *Server) CreatePetition(_ context.Context, req *pb.CreatePetitionRequest) (*pb.CreatePetitionResponse, error) {
	petition := req.GetPetition()

	newPetition := models.Petition{
		Title:       petition.Title,
		Category:    petition.Category,
		Description: petition.Description,
		Image:       petition.Image,
		Status:      uint(petition.Status),
		UserID:      uint(petition.UserId),
	}

	savedPetitionID, err := s.Svc.CreateNew(newPetition)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePetitionResponse{
		PetitionId: uint32(savedPetitionID),
	}, nil
}

func (s *Server) GetPetitions(_ context.Context, req *pb.GetPetitionsRequest) (*pb.GetPetitionsResponse, error) {
	pag := util.Pagination{
		Page:  int(req.Page),
		Limit: int(req.Limit),
	}

	petitions := s.Svc.GetAll(pag)

	getPetitionsResponse := make([]*pb.Petition, len(petitions))

	for i := range getPetitionsResponse {
		p := petitions[i]
		getPetitionsResponse[i] = &pb.Petition{
			PetitionId:  uint32(p.ID),
			Title:       p.Title,
			Category:    p.Category,
			Description: p.Description,
			Image:       p.Image,
			Status:      uint32(p.Status),
			UserId:      uint32(p.UserID),
		}
	}
	return &pb.GetPetitionsResponse{
		Petitions: getPetitionsResponse,
	}, nil
}

func (controller *Server) Update(ctx context.Context, req *pb.UpdatePetitionRequest) (*pb.UpdatePetitionResponse, error) {
	id := req.Id
	status := req.Status

	err := controller.Svc.Update(id, status)
	if err != nil {
		return nil, err
	}
	response := &pb.UpdatePetitionResponse{Message: "Petition status updated successfully"}
	return response, nil
}

func (controller *Server) Delete(ctc context.Context, req *pb.DeletePetitionRequest) (*pb.DeletePetitionResponse, error) {
	id := req.Id
	err := controller.Svc.Delete(id)
	if err != nil {
		return nil, err
	}
	response := &pb.DeletePetitionResponse{Message: "Petition deleted successfully"}
	return response, nil
}
