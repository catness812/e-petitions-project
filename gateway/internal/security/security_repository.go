package security

import (
	"context"
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/security/pb"
	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gookit/slog"
)

type SecurityRepository struct {
	cfg    *config.Config
	client pb.SecurityServiceClient
}

func NewSecurityRepository(cfg *config.Config, client pb.SecurityServiceClient) *SecurityRepository {
	return &SecurityRepository{cfg: cfg, client: client}
}

func (repo *SecurityRepository) Login(loginUser model.UserCredentials) (model.Tokens, error) {

	res, err := repo.client.Login(context.Background(), &pb.UserCredentials{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	})
	var tokens model.Tokens
	if err != nil {
		slog.Errorf("Could not login user: %v", err)
		return tokens, err
	}
	tokens.AccessToken = res.AccessToken
	tokens.RefreshToken = res.RefreshToken
	return tokens, nil
}

func (repo *SecurityRepository) Refresh(token string) (model.Tokens, error) {
	res, err := repo.client.RefreshSession(context.Background(), &pb.RefreshRequest{
		Token: token,
	})
	if err != nil {
		slog.Errorf("Could not refresh user session: %v", err)
		return model.Tokens{}, err
	}
	tokens := model.Tokens{
		AccessToken:  res.Tokens["access_token"],
		RefreshToken: res.Tokens["refresh_token"],
	}
	return tokens, nil
}

func (repo *SecurityRepository) SendOTP(email string) (string, error) {
	res, err := repo.client.SendOTP(context.Background(), &pb.OTPInfo{Email: email})
	if err != nil {
		slog.Errorf("Could not send OTP: %v", err)
		return "", err
	}
	return res.Email, nil
}

func (repo *SecurityRepository) ValidateOTP(otp, email string) (bool, error) {
	validated, err := repo.client.ValidateOTP(context.Background(), &pb.OTPInfo{OTP: otp, Email: email})
	if err != nil {
		slog.Errorf("Could not validate OTP: %v", err)
		return false, err
	}
	return validated.Validated, nil
}
