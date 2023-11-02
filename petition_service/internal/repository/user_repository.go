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
	rpcClient pb.UserServiceClient
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
		return "", err
	}

	return res.Message, nil
}

func (userRepo *UserRepository) CheckUserExistence(id uint) (bool, error) {
	res, err := userRepo.rpcClient.CheckUserExistence(context.Background(),
		&pb.CheckUserExistenceRequest{Id: uint32(id)},
	)
	if err != nil {
		return false, err
	}
	return res.Message, nil
}

func NewUserControllerClient() pb.UserServiceClient {
	cc, err := grpc.Dial(config.Cfg.UserService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Fatal("Could not connect:", err)
	}

	client := pb.NewUserServiceClient(cc)

	return client
}
