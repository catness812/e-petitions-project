package service

import (
	"fmt"
	"time"

	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"

	"github.com/gookit/slog"
	"github.com/robfig/cron/v3"
)

type IPetitionRepository interface {
	Save(petition *models.Petition) error
	GetAll(pagination util.Pagination) []models.Petition
	GetPetitionsByStatus(status models.Status, pagination util.Pagination) ([]models.Petition, error)
	UpdateStatus(id uint, statusID uint) error
	Delete(id uint) error
	GetStatusByTitle(title string) (models.Status, error)
	GetByID(id uint) (models.Petition, error)
	GetAllUserPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error)
	SaveVote(Vote *models.Vote) error
	CheckIfExists(id uint) error
	GetAllUserVotedPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error)
	UpdateCurrVotes(petition models.Petition) error
	HasUserVoted(userID, petitionID uint) error
	GetPetitionsTitles(pagination util.Pagination) ([]models.PetitionInfo, error)
	SearchPetitionsByTitle(searchTerm string, pagination util.Pagination) ([]models.PetitionInfo, error)
}

type IPublisherRepository interface {
	PublishMessage(email string, message string) error
}

type IUserRepository interface {
	GetEmailById(id uint) (string, error)
	CheckUserExistence(id uint) (bool, error)
}

type PetitonService struct {
	petitionRepository  IPetitionRepository
	publisherRepository IPublisherRepository
	userRepository      IUserRepository
}

func NewPetitionService(
	petRepo IPetitionRepository,
	pubRepo IPublisherRepository,
	userRepo IUserRepository,
) *PetitonService {
	return &PetitonService{
		petitionRepository:  petRepo,
		publisherRepository: pubRepo,
		userRepository:      userRepo,
	}
}

func (svc *PetitonService) CreateNew(petition models.Petition) (uint, error) {
	slog.Infof("Creating petition %s", petition.Title)
	// save with draft status when created
	status, err := svc.petitionRepository.GetStatusByTitle(models.DRAFT)
	if err != nil {
		slog.Errorf("Could not retrieve the DRAFT status: %s", err)
		return 0, err
	}
	petition.Status = status

	// get user's email from User Service
	email, err := svc.userRepository.GetEmailById(petition.UserID)
	if err != nil {
		slog.Errorf("Could not retrieve the email from User Service: %s", err)
		return 0, err
	}
	err = svc.publisherRepository.PublishMessage(email, fmt.Sprintf(`Petition "%s" has been successfully created!`, petition.Title))
	if err != nil {
		slog.Errorf("Could not publish message: %s", err)
		return 0, err
	}

	if err := svc.petitionRepository.Save(&petition); err != nil {
		slog.Errorf("Could not create the petition: %s", err)
		return 0, err
	}

	slog.Infof("Petition %s created successfully", petition.Title)
	return petition.ID, nil
}

func (svc *PetitonService) CreateVote(vote models.Vote) error {
	slog.Infof("Creating vote, user %d, petition %d", vote.UserID, vote.PetitionID)
	// see if petition exists
	petition, err := svc.petitionRepository.GetByID(vote.PetitionID)
	if err != nil {
		slog.Errorf("Could not create vote: %s", err)
		return err
	}

	// check if user exists
	exists, err := svc.userRepository.CheckUserExistence(vote.UserID)
	if err != nil {
		slog.Errorf("Could not check if user exists: %s", err)
		return err
	}
	if !exists {
		slog.Errorf("Could not create vote - user %d does not exist: %s", vote.UserID, err)
		return fmt.Errorf("user doesn't exists")
	}

	if err := svc.petitionRepository.HasUserVoted(vote.UserID, vote.PetitionID); err != nil {
		slog.Errorf("Could not check if user %d has voted petition %d: %s", vote.UserID, vote.PetitionID, err)
		return err
	}

	petition.CurrVotes++
	if err := svc.petitionRepository.UpdateCurrVotes(petition); err != nil {
		slog.Errorf("Could not update current votes for petition %d: %s", petition.ID, err)
		return err
	}

	if err := svc.petitionRepository.SaveVote(&vote); err != nil {
		slog.Errorf("Could not save vote: %s", err)
		return err
	}

	// get user's email from User Service
	email, err := svc.userRepository.GetEmailById(petition.UserID)
	if err != nil {
		slog.Errorf("Could not retrieve the email from User Service: %s", err)
		return err
	}

	// if the vote & petition were saved successfully, send email on vote goal
	if petition.VoteGoal == petition.CurrVotes {
		err = svc.publisherRepository.PublishMessage(email, fmt.Sprintf(
			`Petition "%s" has been reached its goal of %d votes! Congrats!`,
			petition.Title, petition.VoteGoal))
		if err != nil {
			slog.Errorf("Could not publish message: %s", err)
			return err
		}
	}

	slog.Info("Vote for user %d, petition %d successfully created", vote.UserID, vote.PetitionID)
	return nil
}

func (svc *PetitonService) GetAll(pagination util.Pagination) []models.Petition {
	slog.Info("Geting all petitions")
	return svc.petitionRepository.GetAll(pagination)
}

func (svc *PetitonService) UpdateStatus(id uint, status string) error {
	slog.Infof("Updating petition %d with status %s", id, status)
	// check if status exists first
	newStatus, err := svc.petitionRepository.GetStatusByTitle(status)
	if err != nil {
		slog.Errorf("Could not retrieve status: %s", err)
		return err
	}
	if err := svc.petitionRepository.UpdateStatus(id, newStatus.ID); err != nil {
		slog.Errorf("Could not update status: %s", err)
		return err
	}

	slog.Infof("Petition %d updated with status %s successfully", id, status)
	return nil
}

func (svc *PetitonService) Delete(id uint) error {
	slog.Infof("Deleting petition %d", id)

	err := svc.petitionRepository.Delete(id)
	if err != nil {
		slog.Errorf("Error deleting petition: %s", err)
		return err
	}

	slog.Infof("Successfully deleted petition %d", id)
	return nil
}

func (svc *PetitonService) GetByID(id uint) (models.Petition, error) {
	slog.Infof("Getting petition %d", id)
	petition, err := svc.petitionRepository.GetByID(id)
	if err != nil {
		slog.Errorf("Error getting petition: %s", err)
		return petition, err
	}

	slog.Infof("Successfully retrieved petition %d", id)
	return petition, nil
}

func (svc *PetitonService) GetAllUserPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error) {
	slog.Info("Geting all petitions for user %d", userID)
	return svc.petitionRepository.GetAllUserPetitions(userID, pagination)
}

func (svc *PetitonService) GetAllUserVotedPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error) {
	slog.Info("Geting all voted petitions for user %d", userID)
	return svc.petitionRepository.GetAllUserVotedPetitions(userID, pagination)
}

func (svc *PetitonService) CheckPetitionExpiration(petition models.Petition) (string, error) {
	slog.Info("Checking if petition %d has expired", petition.ID)
	if time.Now().After(petition.ExpDate) {
		slog.Info("Petition %d has expired", petition.ID)
		email, err := svc.userRepository.GetEmailById(petition.UserID)
		if err != nil {
			slog.Errorf("Could not retrieve email: %s", err)
			return "", err
		}

		err = svc.UpdateStatus(petition.ID, "ARCHIVE")
		if err != nil {
			slog.Errorf("Could not update status: %s", err)
			return "", err
		}

		err = svc.publisherRepository.PublishMessage(email, fmt.Sprintf(`Petition "%s" has expired! It's been moved to your archived petitions.`, petition.Title))
		if err != nil {
			slog.Errorf("Could not publish email: %s", err)
			return "", err
		}
	}
	slog.Info("Petition %d has NOT expired", petition.ID)
	return "", nil
}

func (svc *PetitonService) ScheduleDailyCheck() {
	c := cron.New()
	slog.Info("Scheduled Expiration Checker successfully started...")
	_, err := c.AddFunc("0 0 * * *", func() {
		resultChan := make(chan struct {
			ID    uint
			Error error
		})
		offset := 0
		limit := 100
		for {
			pag := util.Pagination{
				Page:  int(offset),
				Limit: int(limit),
			}
			petitions, err := svc.GetAllActive(pag)
			if err != nil {
				slog.Error(err)
				return
			}

			if len(petitions) == 0 {
				slog.Info("No active petitions found for now...")
			}

			for _, petition := range petitions {
				go func(petition models.Petition) {
					_, err := svc.CheckPetitionExpiration(petition)
					resultChan <- struct {
						ID    uint
						Error error
					}{
						ID:    petition.ID,
						Error: err,
					}
				}(petition)
			}

			for range petitions {
				result := <-resultChan
				if result.Error != nil {
					slog.Infof("Error checking expiration for petition %v: %v", result.ID, result.Error)
				}
			}

			if len(petitions) == 0 {
				break
			}
			offset += limit
		}
	})

	if err != nil {
		slog.Fatalf("Failed to add cron job: %v", err)
	}

	c.Start()

	select {}
}

func (svc *PetitonService) GetAllActive(pagination util.Pagination) ([]models.Petition, error) {
	slog.Infof("Getting all active petitions")
	status, err := svc.petitionRepository.GetStatusByTitle(models.PUBLIC)
	if err != nil {
		slog.Errorf("Could not retrieve status: %s", err)
		return nil, err
	}
	petitions, err := svc.petitionRepository.GetPetitionsByStatus(status, pagination)
	if err != nil {
		slog.Errorf("Could not retrieve petitions: %s", err)
		return nil, err
	}

	slog.Info("Successfully got active petitions")
	return petitions, nil
}

func (svc *PetitonService) GetAllSimilarPetitions(title string) ([]models.PetitionInfo, error) {
	slog.Infof("Getting all similar petitions to %s", title)
	offset := 0
	limit := 100
	similarPetitions := make([]models.PetitionInfo, 0)
	for {
		pag := util.Pagination{
			Page:  offset,
			Limit: limit,
		}

		petitions, err := svc.petitionRepository.GetPetitionsTitles(pag)
		if err != nil {
			slog.Errorf("Error getting petitions title: %s", err)
			return nil, err
		}
		if len(petitions) == 0 {
			break
		}
		processedTitle := util.PreprocessText(title)
		for _, petition := range petitions {
			similarity := util.CalculateTitleSimilarity(processedTitle, petition.Title)
			if similarity >= 0.5 {
				similarPetitions = append(similarPetitions, petition)

			}
		}
		offset += limit
	}
	return similarPetitions, nil

}

func (svc *PetitonService) SearchPetitionsByTitle(searchTerm string, pagination util.Pagination) ([]models.PetitionInfo, error) {
	slog.Infof("Searching petitions with term %s", searchTerm)
	similarPetitions, err := svc.petitionRepository.SearchPetitionsByTitle(searchTerm, pagination)
	if err != nil {
		slog.Errorf("Error searching petitions: %s", err)
		return nil, err
	}
	return similarPetitions, nil
}
