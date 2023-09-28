package security

type ISecurityService interface {
	Login()
	Refresh()
}

func NewSecurityService(repo ISecurityRepository) (ISecurityService, error) {
	return &securityService{
		repo: repo,
	}, nil
}

type securityService struct {
	repo ISecurityRepository
}

func (svc *securityService) Login() {

}

func (svc *securityService) Refresh() {

}
