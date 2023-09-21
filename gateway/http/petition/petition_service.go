package petition

import (
	"log"

	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/catness812/e-petitions-project/gateway/http/petition/pb"
	"google.golang.org/grpc"
)

type IPetitionService interface {
	Get()
	Delete()
	Create()
	Update()
	GetAll()
	Sign()
}

func NewPetitionService(c config.Config) (IPetitionService, error) {
	conn, err := grpc.Dial(c.UserPort, grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect:", err)
	}
	client := pb.NewPetitionServiceClient(conn)

	us := &petitionService{
		cfg:    c,
		conn:   conn,
		client: client,
	}

	return us, nil
}

type petitionService struct {
	cfg    config.Config
	conn   *grpc.ClientConn
	client pb.PetitionServiceClient
}

func (svc *petitionService) Get() {

}

func (svc *petitionService) Delete() {

}

func (svc *petitionService) Create() {

}

func (svc *petitionService) Update() {

}

func (svc *petitionService) GetAll() {

}

func (svc *petitionService) Sign() {

}
