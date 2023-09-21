package petition

import (
	"log"

	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/catness812/e-petitions-project/gateway/http/petition/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.PetitionServiceClient
}

func InitPetitionServiceClient(c *config.Config) pb.PetitionServiceClient {
	cc, err := grpc.Dial(c.PetitionPort, grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect:", err)
	}
	return pb.NewPetitionServiceClient(cc)
}

// func (s *petitionService) NewPetition() error {
// 	return nil
// }

// func (s *petitionService) Update() error {
// 	return nil
// }
// func (s *petitionService) Delete() error {
// 	return nil
// }
// func (s *petitionService) SignPetition() error {
// 	return nil
// }
// func (s *petitionService) GetPetition() error {
// 	return nil
// }
// func (s *petitionService) UpdatePetitionStatus() error {
// 	return nil
// }
