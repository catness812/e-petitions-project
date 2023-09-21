package user

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/catness812/e-petitions-project/gateway/http/user/pb"
	"google.golang.org/grpc"
	"log"
)

type ServiceClient struct {
	Client pb.UserControllerClient
}

func InitUserControllerClient(c *config.Config) pb.UserControllerClient {
	cc, err := grpc.Dial(c.UserPort, grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect:", err)
	}

	return pb.NewUserControllerClient(cc)
}
