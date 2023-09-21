package petition

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/catness812/e-petitions-project/gateway/http/petition/pb"
	"google.golang.org/grpc"
)

type IPetitionService interface {
}

type petitionService struct {
	cfg    config.Config
	conn   *grpc.ClientConn
	client pb.PetitionServiceClient
}

//func (s *petitionService) Delete(id string) error {
//	_, err := s.client.Delete(context.Background(), &pb.{
//		PId: pid,
//	})
//	return err
//}
