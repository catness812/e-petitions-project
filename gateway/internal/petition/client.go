package petition

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition/pb"
	"google.golang.org/grpc"
	"log"
)

type PetitionClient struct {
	Client pb.PetitionServiceClient
}

func InitPetitonServiceClient(c config.Config) pb.PetitionServiceClient {
	cc, err := grpc.Dial(c.PetitionPort, grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect:", err)
	}

	return pb.NewPetitionServiceClient(cc)
}
