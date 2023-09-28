package security

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/security/pb"
)

type ISecurityRepository interface {
	Login()
	Refresh()
}

func NewSecurityRepository(c config.Config, client pb.SecurityServiceClient) (ISecurityRepository, error) {

	us := &securityRepository{
		cfg:    c,
		client: client,
	}

	return us, nil
}

type securityRepository struct {
	cfg    config.Config
	client pb.SecurityServiceClient
}

func (svc *securityRepository) Login() {

}

func (svc *securityRepository) Refresh() {

}
