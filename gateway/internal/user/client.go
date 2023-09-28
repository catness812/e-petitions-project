package user

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"google.golang.org/grpc"
	"log"
)

type UserClient struct {
	Client pb.UserControllerClient
}

func InitUserServiceClient(c config.Config) pb.UserControllerClient {
	cc, err := grpc.Dial(c.UserPort, grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect:", err)
	}

	return pb.NewUserControllerClient(cc)
}
