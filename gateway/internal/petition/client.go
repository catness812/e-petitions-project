package petition

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition/pb"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PetitionClient struct {
	Client pb.PetitionServiceClient
}

func InitPetitonServiceClient(c *config.Config) pb.PetitionServiceClient {
	cc, err := grpc.Dial(c.PetitionPort, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Fatalf("Could not connect: %v", err)
	}

	return pb.NewPetitionServiceClient(cc)
}
