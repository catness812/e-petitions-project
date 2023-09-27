package user

type IUserService interface {
	Get()
	Delete()
	Create()
	Update()
}

func NewUserService(repo IUserRepository) (IUserService, error) {
	return &userService{
		repo: repo,
	}, nil
}

type userService struct {
	repo IUserRepository
}

func (svc *userService) Get() {
}

func (svc *userService) Delete() {

}

func (svc *userService) Create() {

}

func (svc *userService) Update() {

}
