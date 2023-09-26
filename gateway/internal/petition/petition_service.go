package petition

type IPetitionService interface {
	Get()
	Delete()
	Create()
	Update()
	GetAll()
	Sign()
}

func NewPetitionService(repo IPetitionRepository) (IPetitionService, error) {
	return &petitionService{
		repo: repo,
	}, nil
}

type petitionService struct {
	repo IPetitionRepository
}

func (svc *petitionService) Get() {

}

func (svc *petitionService) Delete() {

}

func (svc *petitionService) Create() {

}

func (svc *petitionService) Update() {

}

func (svc *petitionService) GetAll() {

}

func (svc *petitionService) Sign() {

}
