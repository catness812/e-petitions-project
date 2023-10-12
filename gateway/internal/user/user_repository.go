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
	Update(createUser model.UserCredentials) (string, error)
	AddAdmin(email string) (string, error)
}

func NewUserRepository(c *config.Config, client pb.UserControllerClient) (IUserRepository, error) {

	us := &userRepository{
		cfg:    c,
		client: client,
	}

	return us, nil
}

type userRepository struct {
	cfg    *config.Config
	client pb.UserControllerClient
}

func (repo *userRepository) GetByEmail(email string) (model.User, error) {
	res, err := repo.client.GetUserByEmail(context.Background(), &pb.GetUserByEmailRequest{
		Email: email,
	})
	if err != nil {
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
		slog.Errorf("Error getting user by id: ", err)
		return res.Value, err
	}
	if res == nil && res.Value == "" {
		slog.Errorf("Response is empty", err)
		return res.Value, error(errors.New("Response is empty"))
	}

	return res.Value, nil
}

func (repo *userRepository) Delete(email string) (string, error) {
	res, err := repo.client.DeleteUser(context.Background(), &pb.DeleteUserRequest{
		Email: email,
	})
	if err != nil {
		return "", err
	}

	if res != nil && res.Value != "" {
		return res.Value, nil
	}

	return "", errors.New("DeleteUser response is empty")
}

func (repo *userRepository) Create(createUser model.UserCredentials) (string, error) {

	res, err := repo.client.CreateUser(context.Background(), &pb.UserRequest{
		Email:    createUser.Email,
		Password: createUser.Password,
	})

	if err != nil {
		return "", err
	}

	if res != nil && res.Value != "" {
		return res.Value, nil
	}

	return "", errors.New("CreateUser response is empty")

}

func (repo *userRepository) Update(createUser model.UserCredentials) (string, error) {
	res, err := repo.client.UpdateUser(context.Background(), &pb.UserRequest{
		Email:    createUser.Email,
		Password: createUser.Password,
	})

	if err != nil {
		return "", err
	}

	if res != nil && res.Value != "" {
		return res.Value, nil
	}

	return "", errors.New("UpdateUser response is empty")
}

func (repo *userRepository) AddAdmin(email string) (string, error) {
	res, err := repo.client.AddAdmin(context.Background(), &pb.AddAdminRequest{
		Email: email,
	})
	if err != nil {
		return "", err
	}

	if res != nil && res.Value != "" {
		return res.Value, nil
	}

	return "", errors.New("AddAdmin response is empty")
}
