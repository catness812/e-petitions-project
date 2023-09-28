package user

import (
	"fmt"
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type UserClient struct {
	Client pb.UserControllerClient
}

func InitUserServiceClient(c *config.Config) pb.UserControllerClient {
	fmt.Println(c.UserPort)
	cc, err := grpc.Dial(c.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Could not connect:", err)
	}

	return pb.NewUserControllerClient(cc)
}
