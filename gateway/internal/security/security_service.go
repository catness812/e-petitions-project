package security

import "github.com/catness812/e-petitions-project/gateway/model"

type ISecurityRepository interface {
	Login(user model.UserCredentials) (model.Tokens, error)
	Refresh(token string) (model.Tokens, error)
	SendOTP(email string) (string, error)
	ValidateOTP(otp, email string) (bool, error)
}

type SecurityService struct {
	repo ISecurityRepository
}

func NewSecurityService(repo ISecurityRepository) *SecurityService {
	return &SecurityService{repo: repo}
}

func (svc *SecurityService) Login(loginUser model.UserCredentials) (model.Tokens, error) {
	return svc.repo.Login(loginUser)

}

func (svc *SecurityService) Refresh(token string) (model.Tokens, error) {
	return svc.repo.Refresh(token)
}

func (svc *SecurityService) SendOTP(email string) (string, error) {
	return svc.repo.SendOTP(email)
}

func (svc *SecurityService) ValidateOTP(otp, mail string) (bool, error) {
	validated, err := svc.repo.ValidateOTP(otp, mail)
	if err != nil {
		return false, err
	}
	return validated, err
}
