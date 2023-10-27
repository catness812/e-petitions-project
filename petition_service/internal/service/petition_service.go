package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
	"gorm.io/gorm"

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
	UpdatePetition(petition *models.PetitionUpdate) error
}

type IPublisherRepository interface {
	PublishMessage(email string, message string) error
}

type IUserRepository interface {
	GetEmailById(id uint) (string, error)
	CheckUserExistence(id uint) (bool, error)
}

type PetitionService struct {
	petitionRepository  IPetitionRepository
	publisherRepository IPublisherRepository
	userRepository      IUserRepository
}

func NewPetitionService(
	petRepo IPetitionRepository,
	pubRepo IPublisherRepository,
	userRepo IUserRepository,
) *PetitionService {
	return &PetitionService{
		petitionRepository:  petRepo,
		publisherRepository: pubRepo,
		userRepository:      userRepo,
	}
}

func (svc *PetitionService) CreateNew(petition models.Petition) (uint, error) {
	// save with draft status when created
	status, err := svc.petitionRepository.GetStatusByTitle(models.DRAFT)
	if err != nil {
		return 0, err
	}
	petition.Status = status
	// get user's email from User Service
	email, err := svc.userRepository.GetEmailById(petition.UserID)
	if err != nil {
		return 0, err
	}
	parts := strings.Split(email, "@")
	nameParts := strings.Split(parts[0], ".")
	petition.AuthorName = strings.Join(nameParts, " ")
	err = svc.publisherRepository.PublishMessage(email, fmt.Sprintf(`Petition "%s" has been successfully created!`, petition.Title))
	if err != nil {
		return 0, err
	}

	if err := svc.petitionRepository.Save(&petition); err != nil {
		return 0, err
	}

	return petition.ID, nil
}

func (svc *PetitionService) CreateVote(vote models.Vote) error {
	// see if petition exists
	petition, err := svc.petitionRepository.GetByID(vote.PetitionID)
	if err != nil {
		return err
	}

	// check if user exists
	exists, err := svc.userRepository.CheckUserExistence(vote.UserID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user doesn't exists")
	}

	if err := svc.petitionRepository.HasUserVoted(vote.UserID, vote.PetitionID); err != nil {
		return err
	}

	petition.CurrVotes++
	if err := svc.petitionRepository.UpdateCurrVotes(petition); err != nil {
		return err
	}

	if err := svc.petitionRepository.SaveVote(&vote); err != nil {
		return err
	}

	// get user's email from User Service
	email, err := svc.userRepository.GetEmailById(petition.UserID)
	if err != nil {
		return err
	}

	// if the vote & petition were saved successfully, send email on vote goal
	if petition.VoteGoal == petition.CurrVotes {
		err = svc.publisherRepository.PublishMessage(email, fmt.Sprintf(
			`Petition "%s" has been reached its goal of %d votes! Congrats!`,
			petition.Title, petition.VoteGoal))
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *PetitionService) GetAll(pagination util.Pagination) []models.Petition {
	return svc.petitionRepository.GetAll(pagination)
}

func (svc *PetitionService) UpdateStatus(id uint, status string) error {
	// check if status exists first
	newStatus, err := svc.petitionRepository.GetStatusByTitle(status)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			slog.Errorf("Status %v not found", status)
			return fmt.Errorf("Status %v not found", status)
		}
		slog.Errorf("error getting status %v", err)
		return fmt.Errorf("error getting status %v", err)
	}
	// update status
	if err := svc.petitionRepository.UpdateStatus(id, newStatus.ID); err != nil {
		if err == gorm.ErrRecordNotFound {
			slog.Errorf("Petition %v not found", id)
			return fmt.Errorf("Petition %v not found", id)
		}
		slog.Errorf("error updating status %v", err)
		return fmt.Errorf("error updating status %v", err)
	}
	slog.Infof("Petition %v status successfully updated", id)
	return nil
}

func (svc *PetitionService) UpdatePetition(petition *models.PetitionUpdate) error {

	if err := svc.petitionRepository.UpdatePetition(petition); err != nil {
		slog.Errorf("Error updating petition: %v", err)
		return err
	}
	slog.Info("Petition updated successfully")
	return nil
}

func (svc *PetitionService) Delete(id uint) error {
	err := svc.petitionRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *PetitionService) GetByID(id uint) (models.Petition, error) {
	petition, err := svc.petitionRepository.GetByID(id)
	if err != nil {
		return petition, err
	}
	return petition, nil
}

func (svc *PetitionService) GetAllUserPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error) {
	return svc.petitionRepository.GetAllUserPetitions(userID, pagination)
}

func (svc *PetitionService) GetAllUserVotedPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error) {
	return svc.petitionRepository.GetAllUserVotedPetitions(userID, pagination)
}

func (svc *PetitionService) CheckPetitionExpiration(petition models.Petition) (string, error) {
	if time.Now().After(petition.ExpDate) {
		email, err := svc.userRepository.GetEmailById(petition.UserID)
		if err != nil {
			return "", err
		}

		err = svc.UpdateStatus(petition.ID, "ARCHIVE")
		if err != nil {
			return "", err
		}

		err = svc.publisherRepository.PublishMessage(email, fmt.Sprintf(`Petition "%s" has expired! It's been moved to your archived petitions.`, petition.Title))
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

func (svc *PetitionService) ScheduleDailyCheck() {
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
				slog.Println("No active petitions found for now...")
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
					slog.Printf("Error checking expiration for petition %v: %v", result.ID, result.Error)
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

func (svc *PetitionService) GetAllActive(pagination util.Pagination) ([]models.Petition, error) {
	status, err := svc.petitionRepository.GetStatusByTitle("PUBLIC")
	if err != nil {
		return nil, err
	}
	petitions, err := svc.petitionRepository.GetPetitionsByStatus(status, pagination)
	if err != nil {
		return nil, err
	}
	return petitions, nil
}

func (svc *PetitionService) GetAllSimilarPetitions(title string) ([]models.PetitionInfo, error) {
	offset := 0
	limit := 100
	similarPetitions := make([]models.PetitionInfo, 0)
	for {
		pag := util.Pagination{
			Page:  int(offset),
			Limit: int(limit),
		}

		petitions, err := svc.petitionRepository.GetPetitionsTitles(pag)
		if err != nil {
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

func (svc *PetitionService) SearchPetitionsByTitle(searchTerm string, pagination util.Pagination) ([]models.PetitionInfo, error) {
	similarPetitions, err := svc.petitionRepository.SearchPetitionsByTitle(searchTerm, pagination)
	if err != nil {
		return nil, err
	}
	return similarPetitions, nil
}
