package repository

import (
	"context"
	"github.com/catness812/e-petitions-project/petition_service/internal/config"
	"github.com/catness812/e-petitions-project/petition_service/internal/pb"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserRepository struct {
	rpcClient pb.UserControllerClient
}

func NewUserRepository() *UserRepository {
	slog.Info("Creating new User Repository...")
	return &UserRepository{
		rpcClient: NewUserControllerClient(),
	}
}

func (userRepo *UserRepository) GetEmailById(id uint) (string, error) {
	res, err := userRepo.rpcClient.GetUserEmailById(context.Background(),
		&pb.GetUserEmailByIdRequest{Id: uint32(id)},
	)

	if err != nil {
		slog.Error(err)
		return "", err
	}

	return res.Value, nil
}

// TODO move this
func NewUserControllerClient() pb.UserControllerClient {
	slog.Info("Connecting to User Service gRPC Server...")
	cc, err := grpc.Dial(config.Cfg.UserService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Fatal("Could not connect:", err)
	}

	client := pb.NewUserControllerClient(cc)

	return client
}
