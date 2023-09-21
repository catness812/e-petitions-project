package security_controller

import (
	"context"

	"github.com/catness812/e-petitions-project/internal/security/pb"
	"github.com/catness812/e-petitions-project/internal/model"
	"github.com/catness812/e-petitions-project/pkg/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

type ISecurityService interface {
	//RefreshUserToken(token, id string) (map[string]string, error)
	Login(user *models.UserCredentialsModel) (map[string]string, error)
	Register(user *models.UserModel) error
	RefreshUserToken(token string, id uint) (map[string]string, error)
}

type SecurityRpcServer struct {
	pb.UnimplementedSecurityServiceServer
	securitySvc ISecurityService
}

func NewSecurityRpcServer(securitySvc ISecurityService) *SecurityRpcServer {
	return &SecurityRpcServer{securitySvc: securitySvc}
}

func (s *SecurityRpcServer) Login(ctx context.Context, req *pb.UserCredentials) (*pb.Tokens, error) {

	userLogin := models.UserCredentialsModel{
		Email : req.GetEmail(),
		Password : req.GetPassword(),
	}

	token, err := s.securitySvc.Login(&userLogin)

	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.Tokens{
		AccessToken: token["access_token"],
		RefreshToken: token["refresh_token"],
	}, nil
}

func (s *SecurityRpcServer) Register(ctx context.Context, req *pb.UserInfo) (*pb.CreateUserRequest, error) {

	user := models.UserModel{
		ID: uint(req.GetId()),
		Email : req.GetEmail(),
		Password : req.GetPassword(),
		Role: "user",
	}

	err := s.securitySvc.Register(&user);

	if err != nil {	
		return nil, status.Error(codes.Internal, err.Error())
	}

	userInfo := &pb.UserInfo{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     "user",
	}

	return &pb.CreateUserRequest{
		User: userInfo,
	}, nil
}

func (s *SecurityRpcServer) RefreshSession(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	refToken := req.Token
	claims, err := token.IsTokenValid(refToken)
	if err != nil {
		return &pb.RefreshResponse{
			Tokens:  nil,
			Message: err.Error(),
		}, nil
	}
	uid := claims["userID"]
	tokenMap, err := s.securitySvc.RefreshUserToken(refToken, uid.(uint))
	if err != nil {
		return &pb.RefreshResponse{
			Tokens:  nil,
			Message: err.Error(),
		}, nil
	}
	return &pb.RefreshResponse{Tokens: tokenMap}, nil
}
