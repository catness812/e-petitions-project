package security

import (
	"context"
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/security/pb"
	"github.com/catness812/e-petitions-project/gateway/model"
)

type ISecurityRepository interface {
	Login(user model.UserCredentials) (model.Tokens, error)
	Refresh(token string) (model.Tokens, string, error)
}

func NewSecurityRepository(c *config.Config, client pb.SecurityServiceClient) (ISecurityRepository, error) {

	us := &securityRepository{
		cfg:    c,
		client: client,
	}

	return us, nil
}

type securityRepository struct {
	cfg    *config.Config
	client pb.SecurityServiceClient
}

func (repo *securityRepository) Login(loginUser model.UserCredentials) (model.Tokens, error) {

	res, err := repo.client.Login(context.Background(), &pb.UserCredentials{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	})
	var tokens model.Tokens
	if err != nil {
		return tokens, err
	}
	tokens.AccessToken = res.AccessToken
	tokens.RefreshToken = res.RefreshToken
	return tokens, nil
}

func (repo *securityRepository) Refresh(token string) (model.Tokens, string, error) {
	res, err := repo.client.RefreshSession(context.Background(), &pb.RefreshRequest{
		Token: token,
	})
	if err != nil {
		return model.Tokens{}, res.Message, err
	}
	tokens := model.Tokens{
		AccessToken:  res.Tokens["access_token"],
		RefreshToken: res.Tokens["refresh_token"],
	}
	return tokens, "", nil
}
