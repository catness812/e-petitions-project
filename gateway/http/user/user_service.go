package user

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/catness812/e-petitions-project/gateway/http/user/pb"
	"google.golang.org/grpc"
)

type IUserService interface {
	Get()
	Delete()
	Create()
	Update()
}

func NewUserService(c config.Config) (IUserService, error) {
	conn, err := grpc.Dial(c.UserPort, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}
	client := pb.NewUserControllerClient(conn)

	us := &userService{
		cfg:    c,
		conn:   conn,
		client: client,
	}

	return us, nil
}

type userService struct {
	cfg    config.Config
	conn   *grpc.ClientConn
	client pb.UserControllerClient
}

func (svc *userService) Get() {
}

func (svc *userService) Delete() {

}

func (svc *userService) Create() {

}

func (svc *userService) Update() {

}
