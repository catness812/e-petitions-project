package petition

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/catness812/e-petitions-project/gateway/http/petition/pb"
	"google.golang.org/grpc"
)

type IPetitionService interface {
}

type PetitionService struct {
	cfg    config.Config
	conn   *grpc.ClientConn
	client pb.PetitionServiceClient
}
