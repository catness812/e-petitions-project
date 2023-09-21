package rpctransport

import (
	"context"

	"github.com/catness812/e-petitions-project/internal/models"
	"github.com/catness812/e-petitions-project/internal/pb"
	"github.com/catness812/e-petitions-project/internal/service"
)

type Server struct {
	pb.UnimplementedPetitionServiceServer
}

func (*Server) CreatePetition(ctx context.Context, req *pb.CreatePetitionRequest) (*pb.CreatePetitionResponse, error) {
	petition := req.GetPetition()
	userID := req.GetUserId()

	newPetition := models.Petition{
		Title:       petition.Title,
		Category:    petition.Category,
		Description: petition.Description,
		Image:       petition.Image,
		Status:      uint(petition.Status),
		UserID:      uint(userID),
	}

	savedPetitionID, err := service.CreateNew(newPetition)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePetitionResponse{
		PetitionId: uint32(savedPetitionID),
	}, nil
}
