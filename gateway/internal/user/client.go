package user

import (
	"fmt"

	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	Client pb.UserControllerClient
}

func InitUserServiceClient(c *config.Config) pb.UserControllerClient {
	fmt.Println(c.UserPort)
	cc, err := grpc.Dial(c.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Fatalf("Could not connect: %v", err)
		return nil
	}

	return pb.NewUserControllerClient(cc)
}
