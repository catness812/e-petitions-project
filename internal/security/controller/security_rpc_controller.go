package security_controller

import "github.com/catness812/e-petitions-project/internal/models"

type ISecurityService interface {
	RefreshUserToken(token, id string) (map[string]string, error)
	Login(user *models.UserCredentialsModel) (map[string]string, error)
}

type SecurityRpcServer struct {
	securitySvc ISecurityService
}

func NewSecurityRpcServer(securitySvc ISecurityService) *SecurityRpcServer {
	return &SecurityRpcServer{securitySvc: securitySvc}
}
