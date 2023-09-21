package security

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/catness812/e-petitions-project/gateway/http/user/pb"
	"google.golang.org/grpc"
	"log"
)

type ISecurityService interface {
	Login()
	Refresh()
}

func NewSecurityService(c config.Config) (ISecurityService, error) {
	conn, err := grpc.Dial(c.UserPort, grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect:", err)
	}
	client := pb.NewUserControllerClient(conn)

	us := &securityService{
		cfg:    c,
		conn:   conn,
		client: client,
	}

	return us, nil
}

type securityService struct {
	cfg    config.Config
	conn   *grpc.ClientConn
	client pb.UserControllerClient
}

func (svc *securityService) Login() {

}

func (svc *securityService) Refresh() {

}
