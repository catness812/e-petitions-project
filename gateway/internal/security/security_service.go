package security

import "github.com/catness812/e-petitions-project/gateway/model"

type ISecurityService interface {
	Login(loginUser model.UserCredentials) (model.Tokens, error)
	Refresh(token string) (model.Tokens, error)
}

func NewSecurityService(repo ISecurityRepository) (ISecurityService, error) {
	return &securityService{
		repo: repo,
	}, nil
}

type securityService struct {
	repo ISecurityRepository
}

func (svc *securityService) Login(loginUser model.UserCredentials) (model.Tokens, error) {
	return svc.repo.Login(loginUser)

}

func (svc *securityService) Refresh(token string) (model.Tokens, error) {
	return svc.repo.Refresh(token)
}
