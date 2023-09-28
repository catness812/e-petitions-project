package security

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/security/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type SecurityClient struct {
	Client pb.SecurityServiceClient
}

func InitAuthServiceClient(c *config.Config) pb.SecurityServiceClient {
	cc, err := grpc.Dial(c.SecurityPort, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Could not connect:", err)
	}

	return pb.NewSecurityServiceClient(cc)
}
