package security_controller

import (
	"context"

	"github.com/catness812/e-petitions-project/security_service/internal/security_pb"

	models "github.com/catness812/e-petitions-project/security_service/internal/model"
	"github.com/catness812/e-petitions-project/security_service/pkg/jwtoken"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ISecurityService interface {
	Login(user *models.UserCredentialsModel) (map[string]string, error)
	RefreshUserToken(token string, email string) (map[string]string, error)
}

type SecurityRpcServer struct {
	securitySvc ISecurityService
}

func NewSecurityRpcServer(securitySvc ISecurityService) *SecurityRpcServer {
	return &SecurityRpcServer{securitySvc: securitySvc}
}

func (s *SecurityRpcServer) Login(ctx context.Context, req *security_pb.UserCredentials) (*security_pb.Tokens, error) {
	userLogin := models.UserCredentialsModel{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	token, err := s.securitySvc.Login(&userLogin)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &security_pb.Tokens{
		AccessToken:  token["access_token"],
		RefreshToken: token["refresh_token"],
	}, nil
}

func (s *SecurityRpcServer) RefreshSession(ctx context.Context, req *security_pb.RefreshRequest) (*security_pb.RefreshResponse, error) {
	refToken := req.Token
	userEmail, err := jwtoken.IsTokenValid(refToken)
	if err != nil {
		return &security_pb.RefreshResponse{
			Tokens:  nil,
			Message: err.Error(),
		}, nil
	}
	tokenMap, err := s.securitySvc.RefreshUserToken(refToken, userEmail)
	if err != nil {
		return &security_pb.RefreshResponse{
			Tokens:  nil,
			Message: err.Error(),
		}, nil
	}
	return &security_pb.RefreshResponse{Tokens: tokenMap}, nil
}

func (s *SecurityRpcServer) ValidateToken(ctx context.Context, req *security_pb.Token) (*security_pb.ValidateTokenResponse, error) {
	email, err := jwtoken.IsTokenValid(req.Token)
		if err != nil {
			return nil, status.Error(codes.NotFound, err.Error()) 
		}
	result :=  &security_pb.ValidateTokenResponse{Token:req.Token, Email:email}

	return result, nil
}