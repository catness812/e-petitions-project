package security_controller

import (
	"context"
	"errors"
	"github.com/catness812/e-petitions-project/security_service/internal/config"
	"github.com/catness812/e-petitions-project/security_service/internal/pb"
	"github.com/gookit/slog"
	"github.com/streadway/amqp"

	models "github.com/catness812/e-petitions-project/security_service/internal/model"
	"github.com/catness812/e-petitions-project/security_service/pkg/jwtoken"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ISecurityService interface {
	Login(user *models.UserCredentialsModel) (map[string]string, string, error)
	RefreshUserToken(token string, email string) (map[string]string, error)
	SendOTP(mail string, ch *amqp.Channel, cfg *config.Config) (string, error)
	VerifyOTP(mail string, userOTP string) error
}

type SecurityRpcServer struct {
	securitySvc ISecurityService
	rabbitCh    *amqp.Channel
	cfg         *config.Config
}

func NewSecurityRpcServer(securitySvc ISecurityService, rabbitCh *amqp.Channel, cfg *config.Config) *SecurityRpcServer {
	return &SecurityRpcServer{securitySvc: securitySvc, rabbitCh: rabbitCh, cfg: cfg}
}

func (s *SecurityRpcServer) Login(ctx context.Context, req *pb.UserCredentials) (*pb.Tokens, error) {
	userLogin := models.UserCredentialsModel{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	token, userId, err := s.securitySvc.Login(&userLogin)
	if err != nil {
		slog.Errorf("Failed to login login user: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.Tokens{
		AccessToken:  token["access_token"],
		RefreshToken: token["refresh_token"],
		UserId:       userId,
	}, nil
}

func (s *SecurityRpcServer) RefreshSession(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	refToken := req.Token
	userEmail, err := jwtoken.IsTokenValid(refToken)
	if err != nil {
		slog.Errorf("Failed to validate token: %v", err)
		return nil, status.Error(codes.Unauthenticated, errors.New("failed to refresh user session").Error())
	}
	tokenMap, err := s.securitySvc.RefreshUserToken(refToken, userEmail)
	if err != nil {
		slog.Errorf("Failed to refresh user session: %v", err)
		return nil, status.Error(codes.Internal, errors.New("failed to refresh user session").Error())
	}
	return &pb.RefreshResponse{Tokens: tokenMap}, nil
}

func (s *SecurityRpcServer) ValidateToken(ctx context.Context, req *pb.Token) (*pb.ValidateTokenResponse, error) {
	email, err := jwtoken.IsTokenValid(req.Token)
	if err != nil {
		slog.Errorf("Failed to validate token: %v", err)
		return nil, status.Error(codes.Unauthenticated, errors.New("failed to validate token").Error())
	}
	result := &pb.ValidateTokenResponse{Token: req.Token, Email: email}
	return result, nil
}

func (s *SecurityRpcServer) SendOTP(ctx context.Context, req *pb.OTPInfo) (*pb.OTPInfo, error) {
	otpCode, err := s.securitySvc.SendOTP(req.Email, s.rabbitCh, s.cfg)
	if err != nil {
		slog.Errorf("Failed to send otp: %v", err)
		return nil, status.Error(codes.Internal, errors.New("failed to send otp").Error())
	}
	return &pb.OTPInfo{OTP: otpCode, Email: req.Email}, nil
}

func (s *SecurityRpcServer) ValidateOTP(ctx context.Context, req *pb.OTPInfo) (*pb.IsOTPValidated, error) {
	if err := s.securitySvc.VerifyOTP(req.Email, req.OTP); err != nil {
		slog.Errorf("Failed to validate otp: %v", err)
		return nil, status.Error(codes.InvalidArgument, errors.New("failed to validate otp").Error())
	}
	return &pb.IsOTPValidated{Validated: true}, nil
}
