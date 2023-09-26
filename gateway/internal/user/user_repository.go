package user

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
)

type IUserRepository interface {
	Get()
	Delete()
	Create()
	Update()
}

func NewUserRepository(c config.Config, client pb.UserControllerClient) (IUserRepository, error) {

	us := &userRepository{
		cfg:    c,
		client: client,
	}

	return us, nil
}

type userRepository struct {
	cfg    config.Config
	client pb.UserControllerClient
}

func (svc *userRepository) Get() {
}

func (svc *userRepository) Delete() {

}

func (svc *userRepository) Create() {

}

func (svc *userRepository) Update() {

}
