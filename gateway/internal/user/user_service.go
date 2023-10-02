package user

import "github.com/catness812/e-petitions-project/gateway/model"

type IUserService interface {
	Get(email string) (model.User, error)
	Delete(email string) (string, error)
	Create(createUser model.UserCredentials) (string, error)
	Update(createUser model.UserCredentials) (string, error)
	AddAdmin(email string) (string, error)
}

func NewUserService(repo IUserRepository) (IUserService, error) {
	return &userService{
		repo: repo,
	}, nil
}

type userService struct {
	repo IUserRepository
}

func (svc *userService) Get(email string) (model.User, error) {
	return svc.repo.Get(email)
}

func (svc *userService) Delete(email string) (string, error) {

	return svc.repo.Delete(email)

}

func (svc *userService) Create(createUser model.UserCredentials) (string, error) {
	return svc.repo.Create(createUser)
}

func (svc *userService) Update(createUser model.UserCredentials) (string, error) {
	return svc.repo.Update(createUser)
}

func (svc *userService) AddAdmin(email string) (string, error) {
	return svc.repo.AddAdmin(email)
}
