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
	GetPetitionsByStatus(status models.Status, pagination util.Pagination, order util.PetitionOrder) ([]models.Petition, error)
	UpdateStatus(uuid string, statusID uint) error
	Delete(uuid string) error
	GetStatusByTitle(title string) (models.Status, error)
	GetByID(uuid string) (models.Petition, error)
	GetAllUserPetitions(userUUID string, pagination util.Pagination) ([]models.Petition, error)
	SaveVote(Vote *models.Vote) error
	CheckIfExists(id string) error
	GetAllUserVotedPetitions(userUUID string, pagination util.Pagination) ([]models.Petition, error)
	UpdateCurrVotes(petition models.Petition) error
	HasUserVoted(userUUID, petitionUUID string) error
	GetPetitionsTitles(pagination util.Pagination) ([]models.PetitionInfo, error)
	SearchPetitionsByTitle(searchTerm string, pagination util.Pagination) ([]models.PetitionInfo, error)
	UpdatePetition(petition *models.PetitionUpdate) error
}

type IPublisherRepository interface {
	PublishMessage(email string, message string) error
}

type IUserRepository interface {
	GetEmailById(id string) (string, error)
	CheckUserExistence(id string) (bool, error)
	GetAdminEmails() ([]string, error)
}

type IElasticSearchRepository interface {
	AddPetition(petition models.Petition) error
	SearchPetitionsByTitle(title string, pagination util.Pagination) ([]models.PetitionInfo, error)
}

type PetitionService struct {
	petitionRepository      IPetitionRepository
	publisherRepository     IPublisherRepository
	userRepository          IUserRepository
	elasticSearchRepository IElasticSearchRepository
}

func NewPetitionService(
	petRepo IPetitionRepository,
	pubRepo IPublisherRepository,
	userRepo IUserRepository,
	elasticSearchRepo IElasticSearchRepository,
) *PetitionService {
	svc := &PetitionService{
		petitionRepository:      petRepo,
		publisherRepository:     pubRepo,
		userRepository:          userRepo,
		elasticSearchRepository: elasticSearchRepo,
	}

	go func() {
		svc.scheduleDailyDigest()
	}()

	return svc
}

func (svc *PetitionService) CreateNew(petition models.Petition) (string, error) {
	slog.Infof("Creating petition %s", petition.Title)
	// save with draft status when created
	status, err := svc.petitionRepository.GetStatusByTitle(models.DRAFT)
	if err != nil {
		slog.Errorf("Could not retrieve the DRAFT status: %s", err)
		return "", err
	}
	petition.Status = status
	// get user's email from User Service
	email, err := svc.userRepository.GetEmailById(petition.UserID)
	if err != nil {
		slog.Errorf("Could not retrieve the email from User Service: %s", err)
		return "", err
	}
	parts := strings.Split(email, "@")
	nameParts := strings.Split(parts[0], ".")
	petition.AuthorName = strings.Join(nameParts, " ")
	err = svc.publisherRepository.PublishMessage(email, fmt.Sprintf(`Petition "%s" has been successfully created!`, petition.Title))
	if err != nil {
		slog.Errorf("Could not publish message: %s", err)
		return "", err
	}

	if err := svc.petitionRepository.Save(&petition); err != nil {
		slog.Errorf("Could not create the petition: %s", err)
		return "", err
	}

	if err := svc.elasticSearchRepository.AddPetition(petition); err != nil {
		slog.Errorf("could not add petition to elastic search: %s", err)
		return "", err
	}

	slog.Infof("Petition %s created successfully", petition.Title)
	return petition.UUID, nil
}

func (svc *PetitionService) CreateVote(vote models.Vote) error {
	slog.Infof("Creating vote, user %s, petition %s", vote.UserID, vote.PetitionUUID)
	// see if petition exists
	petition, err := svc.petitionRepository.GetByID(vote.PetitionUUID)
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
		slog.Errorf("Could not create vote - user %s does not exist: %s", vote.UserID, err)
		return fmt.Errorf("user doesn't exists")
	}

	if err := svc.petitionRepository.HasUserVoted(vote.UserID, vote.PetitionUUID); err != nil {
		slog.Errorf("Could not check if user %s has voted petition %d: %s", vote.UserID, vote.PetitionUUID, err)
		return err
	}

	petition.CurrVotes++
	if err := svc.petitionRepository.UpdateCurrVotes(petition); err != nil {
		slog.Errorf("Could not update current votes for petition %s: %s", petition.UUID, err)
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
			slog.Errorf("could not publish message: %s", err)
			return err
		}
	}

	slog.Info("Vote for user %s, petition %s successfully created", vote.UserID, vote.PetitionUUID)
	return nil
}

func (svc *PetitionService) GetAll(pagination util.Pagination) []models.Petition {
	slog.Info("Geting all petitions")
	return svc.petitionRepository.GetAll(pagination)
}

func (svc *PetitionService) UpdateStatus(uuid string, status string) error {
	// check if status exists first
	newStatus, err := svc.petitionRepository.GetStatusByTitle(status)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			slog.Errorf("Status %v not found", status)
			return fmt.Errorf("status %v not found", status)
		}
		slog.Errorf("error getting status %v", err)
		return fmt.Errorf("error getting status %v", err)
	}
	// update status
	if err := svc.petitionRepository.UpdateStatus(uuid, newStatus.ID); err != nil {
		if err == gorm.ErrRecordNotFound {
			slog.Errorf("Petition %v not found", uuid)
			return fmt.Errorf("petition %v not found", uuid)
		}
		slog.Errorf("error updating status %v", err)
		return fmt.Errorf("error updating status %v", err)
	}
	slog.Infof("Petition %v status successfully updated", uuid)
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

func (svc *PetitionService) Delete(uuid string) error {
	slog.Infof("Deleting petition %d", uuid)
	err := svc.petitionRepository.Delete(uuid)
	if err != nil {
		slog.Errorf("Error deleting petition: %s", err)
		return err
	}

	slog.Infof("Successfully deleted petition %d", uuid)
	return nil
}

func (svc *PetitionService) GetByID(uuid string) (models.Petition, error) {
	slog.Infof("Getting petition %s", uuid)
	petition, err := svc.petitionRepository.GetByID(uuid)
	if err != nil {
		slog.Errorf("Error getting petition: %s", err)
		return petition, err
	}

	slog.Infof("Successfully retrieved petition %s", uuid)
	return petition, nil
}

func (svc *PetitionService) GetAllUserPetitions(userUUID string, pagination util.Pagination) ([]models.Petition, error) {
	slog.Info("Geting all petitions for user %s", userUUID)
	return svc.petitionRepository.GetAllUserPetitions(userUUID, pagination)
}

func (svc *PetitionService) GetAllUserVotedPetitions(userUUID string, pagination util.Pagination) ([]models.Petition, error) {
	slog.Info("Geting all voted petitions for user %s", userUUID)
	return svc.petitionRepository.GetAllUserVotedPetitions(userUUID, pagination)
}

func (svc *PetitionService) CheckPetitionExpiration(petition models.Petition) (string, error) {
	slog.Info("Checking if petition %s has expired", petition.UUID)
	if time.Now().After(petition.ExpDate) {
		slog.Info("Petition %d has expired", petition.UUID)
		email, err := svc.userRepository.GetEmailById(petition.UserID)
		if err != nil {
			slog.Errorf("Could not retrieve email: %s", err)
			return "", err
		}

		err = svc.UpdateStatus(petition.UUID, "ARCHIVE")
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
	slog.Info("Petition %s has NOT expired", petition.UUID)
	return "", nil
}

// ScheduleDailyDigest sends a daily digest at 10:00
// of petitions that are IN_REVIEW to admins
func (svc *PetitionService) scheduleDailyDigest() {
	c := cron.New()
	slog.Info("Scheduled Daily Petition Digest successfully started...")
	_, err := c.AddFunc("0 10 * * *", func() {
		slog.Info("Sending petition digest to admins...")
		emails, err := svc.userRepository.GetAdminEmails()
		if err != nil {
			slog.Errorf("Could not retrieve admin emails: %s", err)
		}

		status, err := svc.petitionRepository.GetStatusByTitle(models.IN_REVIEW)
		if err != nil {
			slog.Errorf("Could not retrieve status: %s", err)
		}

		page := 1
		for _, email := range emails {
			petitions, err := svc.petitionRepository.GetPetitionsByStatus(status, util.Pagination{
				Page:  page,
				Limit: 30,
			}, util.PetitionOrder{CreatedAt: util.ASC})
			if err != nil {
				slog.Errorf("Could not retrieve petitions: %s", err)
			}

			// TODO need to make a new mail template
			message := "Here are some petitions awaiting your review!\n"
			for _, pet := range petitions {
				slog.Info(pet.Title)
				message += fmt.Sprintf("* Title: %s; UUID: %s; Created At: %s \n", pet.Title, pet.UUID, pet.CreatedAt)
			}

			if len(petitions) > 0 {
				err := svc.publisherRepository.PublishMessage(email, message)
				if err != nil {
					slog.Errorf("Could not publish message: %s", err)
				}
			}

			page += 1
		}
		slog.Info("Successfully sent digest")
	})

	if err != nil {
		slog.Fatalf("Failed to add cron job: %v", err)
	}

	c.Start()

	select {}
}

func (svc *PetitionService) ScheduleDailyCheck() {
	c := cron.New()
	slog.Info("Scheduled Expiration Checker successfully started...")
	_, err := c.AddFunc("0 0 * * *", func() {
		resultChan := make(chan struct {
			ID    string
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
						ID    string
						Error error
					}{
						ID:    petition.UUID,
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

func (svc *PetitionService) GetAllActive(pagination util.Pagination) ([]models.Petition, error) {
	slog.Infof("Getting all active petitions")
	status, err := svc.petitionRepository.GetStatusByTitle(models.PUBLIC)
	if err != nil {
		slog.Errorf("Could not retrieve status: %s", err)
		return nil, err
	}
	petitions, err := svc.petitionRepository.GetPetitionsByStatus(status, pagination, util.PetitionOrder{})
	if err != nil {
		slog.Errorf("Could not retrieve petitions: %s", err)
		return nil, err
	}

	slog.Info("Successfully got active petitions")
	return petitions, nil
}

func (svc *PetitionService) GetAllSimilarPetitions(title string) ([]models.PetitionInfo, error) {
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

func (svc *PetitionService) SearchPetitionsByTitle(searchTerm string, pagination util.Pagination) ([]models.PetitionInfo, error) {

	slog.Infof("Searching petitions with term %s", searchTerm)
	// similarPetitions, err := svc.petitionRepository.SearchPetitionsByTitle(searchTerm, pagination)
	similarPetitions, err := svc.elasticSearchRepository.SearchPetitionsByTitle(searchTerm, pagination)
	if err != nil {
		slog.Errorf("Error searching petitions: %s", err)
		return nil, err
	}
	return similarPetitions, nil
}
