package user

import (
	"context"
	"errors"
	"time"

	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gookit/slog"
)

type UserRepository struct {
	cfg    *config.Config
	client pb.UserServiceClient
}

func NewUserRepository(cfg *config.Config, client pb.UserServiceClient) *UserRepository {
	return &UserRepository{cfg: cfg, client: client}
}

func (repo *UserRepository) GetByEmail(email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(repo.cfg.LongTimeout)*time.Second)
	defer cancel()

	res, err := repo.client.GetUserByEmail(ctx, &pb.GetUserByEmailRequest{
		Email: email,
	})
	if err != nil {
		slog.Errorf("Error getting user by email: %v", err)
		return model.User{}, err
	}
	var user model.User
	user.UserCredentials.Email = res.Email
	user.UserCredentials.Password = res.Password
	user.UUID = res.Uuid
	user.Role = res.Role

	return user, nil
}

func (repo *UserRepository) GetByID(id string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(repo.cfg.LongTimeout)*time.Second)
	defer cancel()

	res, err := repo.client.GetUserEmailById(ctx, &pb.GetUserEmailByIdRequest{
		Uuid: id,
	})
	if err != nil {
		slog.Errorf("Error getting user by id: %v", err)
		return res.Message, err
	}
	if res == nil && res.Message == "" {
		slog.Error("Response is empty")
		return res.Message, errors.New("Response is empty ")
	}

	return res.Message, nil
}

func (repo *UserRepository) Delete(email string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(repo.cfg.LongTimeout)*time.Second)
	defer cancel()

	res, err := repo.client.DeleteUser(ctx, &pb.DeleteUserRequest{
		Email: email,
	})
	if err != nil {
		slog.Errorf("Error deleting user: %v", err)
		return "", err
	}

	if res == nil && res.Message == "" {
		slog.Error("DeleteUser response is empty")
		return res.Message, errors.New("DeleteUser response is empty")
	}

	return res.Message, nil
}

func (repo *UserRepository) Create(createUser model.UserCredentials) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(repo.cfg.LongTimeout)*time.Second)
	defer cancel()

	res, err := repo.client.CreateUser(ctx, &pb.UserRequest{
		Email:    createUser.Email,
		Password: createUser.Password,
	})

	if err != nil {
		slog.Errorf("Error creating user: %v", err)

		return "", err
	}

	if res == nil && res.Message == "" {
		slog.Error("CreateUser response is empty")
		return res.Message, errors.New("CreateUser response is empty")
	}

	return res.Message, nil

}

func (repo *UserRepository) OTPCreate(createUser model.UserCredentials) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(repo.cfg.LongTimeout)*time.Second)
	defer cancel()

	res, err := repo.client.CreateUserOTP(ctx, &pb.UserRequest{
		Email:    createUser.Email,
		Password: createUser.Password,
	})

	if err != nil {
		slog.Errorf("Error creating OTP user: %v", err)

		return "", err
	}

	if res == nil && res.Message == "" {
		slog.Error("CreateUserOTP response is empty")
		return res.Message, errors.New("CreateUserOTP response is empty")
	}

	return res.Message, nil

}

func (repo *UserRepository) Update(createUser model.UserCredentials) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(repo.cfg.LongTimeout)*time.Second)
	defer cancel()

	res, err := repo.client.UpdateUser(ctx, &pb.UserRequest{
		Email:    createUser.Email,
		Password: createUser.Password,
	})

	if err != nil {
		slog.Errorf("Error updating user: %v", err)
		return "", err
	}

	if res == nil && res.Message == "" {
		slog.Errorf("UpdateUser response is empty")
		return res.Message, nil
	}

	return res.Message, nil
}

func (repo *UserRepository) AddAdmin(email string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(repo.cfg.LongTimeout)*time.Second)
	defer cancel()

	res, err := repo.client.AddAdmin(ctx, &pb.AddAdminRequest{
		Email: email,
	})
	if err != nil {
		slog.Errorf("Error adding admin: %v", err)
		return "", err
	}

	if res == nil && res.Message == "" {
		slog.Errorf("AddAdmin response is empty")
		return res.Message, errors.New("AddAdmin response is empty")
	}

	return res.Message, nil
}
