package security

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/security/pb"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SecurityClient struct {
	Client pb.SecurityServiceClient
}

func InitAuthServiceClient(c *config.Config) pb.SecurityServiceClient {
	cc, err := grpc.Dial(c.SecurityPort, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Fatal("Could not connect: %v", err)
	}

	return pb.NewSecurityServiceClient(cc)
}
