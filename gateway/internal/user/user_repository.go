package user

import (
	"context"
	"errors"

	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gookit/slog"
)

type IUserRepository interface {
	GetByEmail(email string) (model.User, error)
	GetByID(id uint32) (string, error)
	Delete(email string) (string, error)
	Create(createUser model.UserCredentials) (string, error)
	OTPCreate(createUser model.UserCredentials) (string, error)
	Update(createUser model.UserCredentials) (string, error)
	AddAdmin(email string) (string, error)
}

func NewUserRepository(c *config.Config, client pb.UserServiceClient) (IUserRepository, error) {

	us := &userRepository{
		cfg:    c,
		client: client,
	}

	return us, nil
}

type userRepository struct {
	cfg    *config.Config
	client pb.UserServiceClient
}

func (repo *userRepository) GetByEmail(email string) (model.User, error) {
	res, err := repo.client.GetUserByEmail(context.Background(), &pb.GetUserByEmailRequest{
		Email: email,
	})
	if err != nil {
		slog.Errorf("Error getting user by email: %v", err)
		return model.User{}, err
	}
	var user model.User
	user.UserCredentials.Email = res.Email
	user.UserCredentials.Password = res.Password
	user.Id = res.Id
	user.Role = res.Role

	return user, nil
}

func (repo *userRepository) GetByID(id uint32) (string, error) {
	res, err := repo.client.GetUserEmailById(context.Background(), &pb.GetUserEmailByIdRequest{
		Id: id,
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

func (repo *userRepository) Delete(email string) (string, error) {
	res, err := repo.client.DeleteUser(context.Background(), &pb.DeleteUserRequest{
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

func (repo *userRepository) Create(createUser model.UserCredentials) (string, error) {

	res, err := repo.client.CreateUser(context.Background(), &pb.UserRequest{
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

func (repo *userRepository) OTPCreate(createUser model.UserCredentials) (string, error) {

	res, err := repo.client.CreateUserOTP(context.Background(), &pb.UserRequest{
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

func (repo *userRepository) Update(createUser model.UserCredentials) (string, error) {
	res, err := repo.client.UpdateUser(context.Background(), &pb.UserRequest{
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

func (repo *userRepository) AddAdmin(email string) (string, error) {
	res, err := repo.client.AddAdmin(context.Background(), &pb.AddAdminRequest{
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
