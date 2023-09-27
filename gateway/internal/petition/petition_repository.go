package petition

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition/pb"
)

type IPetitionRepository interface {
	Get()
	Delete()
	Create()
	Update()
}

func NewPetitionRepository(c config.Config, client pb.PetitionServiceClient) (IPetitionRepository, error) {

	us := &petitionRepository{
		cfg:    c,
		client: client,
	}

	return us, nil
}

type petitionRepository struct {
	cfg    config.Config
	client pb.PetitionServiceClient
}

func (svc *petitionRepository) Get() {
}

func (svc *petitionRepository) Delete() {

}

func (svc *petitionRepository) Create() {

}

func (svc *petitionRepository) Update() {

}
