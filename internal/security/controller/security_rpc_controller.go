package security_controller

import (
	"context"
	"github.com/catness812/e-petitions-project/internal/model"
	"github.com/catness812/e-petitions-project/internal/security/security_pb"
	"github.com/catness812/e-petitions-project/pkg/jwtoken"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ISecurityService interface {
	Login(user *models.UserCredentialsModel) (map[string]string, error)
	Register(user *models.UserModel) error
	RefreshUserToken(token string, id uint) (map[string]string, error)
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

func (s *SecurityRpcServer) Register(ctx context.Context, req *security_pb.UserInfo) (*security_pb.CreateUserRequest, error) {
	user := models.UserModel{
		ID:       uint(req.GetId()),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     "user",
	}
	err := s.securitySvc.Register(&user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	userInfo := &security_pb.UserInfo{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     "user",
	}
	return &security_pb.CreateUserRequest{
		User: userInfo,
	}, nil
}

func (s *SecurityRpcServer) RefreshSession(ctx context.Context, req *security_pb.RefreshRequest) (*security_pb.RefreshResponse, error) {
	refToken := req.Token
	claims, err := jwtoken.IsTokenValid(refToken)
	if err != nil {
		return &security_pb.RefreshResponse{
			Tokens:  nil,
			Message: err.Error(),
		}, nil
	}
	uid := claims["userID"]
	uid64 := uid.(float64)
	tokenMap, err := s.securitySvc.RefreshUserToken(refToken, uint(uid64))
	if err != nil {
		return &security_pb.RefreshResponse{
			Tokens:  nil,
			Message: err.Error(),
		}, nil
	}
	return &security_pb.RefreshResponse{Tokens: tokenMap}, nil
}
