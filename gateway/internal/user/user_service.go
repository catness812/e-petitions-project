package user

import "github.com/catness812/e-petitions-project/gateway/model"

type IUserRepository interface {
	GetByEmail(email string) (model.User, error)
	GetByID(id uint32) (string, error)
	Delete(email string) (string, error)
	Create(createUser model.UserCredentials) (string, error)
	OTPCreate(createUser model.UserCredentials) (string, error)
	Update(createUser model.UserCredentials) (string, error)
	AddAdmin(email string) (string, error)
}

type UserService struct {
	repo IUserRepository
}

func NewUserService(repo IUserRepository) *UserService {
	return &UserService{repo: repo}
}

func (svc *UserService) GetByEmail(email string) (model.User, error) {
	return svc.repo.GetByEmail(email)
}

func (svc *UserService) GetByID(id uint32) (string, error) {
	return svc.repo.GetByID(id)
}

func (svc *UserService) Delete(email string) (string, error) {
	return svc.repo.Delete(email)
}
func (svc *UserService) Create(createUser model.UserCredentials) (string, error) {
	return svc.repo.Create(createUser)
}
func (svc *UserService) OTPCreate(createUser model.UserCredentials) (string, error) {
	return svc.repo.Create(createUser)
}

func (svc *UserService) Update(createUser model.UserCredentials) (string, error) {
	return svc.repo.Update(createUser)
}

func (svc *UserService) AddAdmin(email string) (string, error) {
	return svc.repo.AddAdmin(email)
}
