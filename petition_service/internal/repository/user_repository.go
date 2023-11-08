package repository

import (
	"context"
	"time"

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

func NewUserControllerClient() pb.UserServiceClient {
	cc, err := grpc.Dial(config.Cfg.UserService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Fatal("Could not connect:", err)
	}

	client := pb.NewUserServiceClient(cc)

	return client
}

func (userRepo *UserRepository) GetEmailById(uuid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := userRepo.rpcClient.GetUserEmailById(ctx,
		&pb.GetUserEmailByIdRequest{Id: uuid},
	)

	if err != nil {
		return "", err
	}

	return res.Message, nil
}

func (userRepo *UserRepository) CheckUserExistence(uuid string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := userRepo.rpcClient.CheckUserExistence(ctx,
		&pb.CheckUserExistenceRequest{Id: uuid},
	)
	if err != nil {
		return false, err
	}
	return res.Message, nil
}

func (userRepo *UserRepository) GetAdminEmails() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := userRepo.rpcClient.GetAdminEmails(ctx, &pb.GetAdminEmailsRequest{})
	if err != nil {
		return nil, err
	}

	return res.AdminEmails, nil
}
