package repository

import (
	"errors"
	"time"

	"github.com/gookit/slog"

	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
	"github.com/catness812/e-petitions-project/petition_service/pkg/database/postgres"
	"gorm.io/gorm"
)

type PetitionRepository struct {
	db *gorm.DB
}

func NewPetitionRepository(db *gorm.DB) *PetitionRepository {
	slog.Info("Creating new Petition Repository...")
	return &PetitionRepository{
		db: db,
	}
}

func (repo *PetitionRepository) Save(petition *models.Petition) error {
	err := repo.db.Create(petition).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PetitionRepository) GetAll(pagination util.Pagination) []models.Petition {
	var petitions []models.Petition
	// find paginated petitions
	repo.db.Scopes(postgres.Paginate(pagination)).Preload("Status").Find(&petitions)
	return petitions
}

func (repo *PetitionRepository) GetPetitionsByStatus(status models.Status, pagination util.Pagination) ([]models.Petition, error) {
	var petitions []models.Petition

	err := repo.db.Preload("Status").Scopes(postgres.Paginate(pagination)).
		Where("status_id = ?", status.ID).Limit(50).Find(&petitions).Error
	if err != nil {
		return nil, err
	}
	return petitions, nil
}

func (repo *PetitionRepository) UpdateStatus(id string, statusID uint) error {
	var petition models.Petition
	err := repo.db.Where("uuid = ?", id).Preload("Status").First(&petition).Error
	if err != nil {
		return err
	}
	petition.Status.ID = statusID
	petition.UpdatedAt = time.Now()

	if err := repo.db.Save(&petition).Error; err != nil {
		return err
	}
	return nil
}

func (repo *PetitionRepository) SaveVote(Vote *models.Vote) error {
	err := repo.db.Create(Vote).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PetitionRepository) UpdatePetition(petition *models.PetitionUpdate) error {
	existingPetition := &models.Petition{}
	slog.Info("petition.id:", petition.UUID)
	err := repo.db.Where("uuid = ?", petition.UUID).First(&existingPetition).Error
	if err != nil {
		return err
	}

	// Update only non-null fields
	if petition.Title != "" {
		existingPetition.Title = petition.Title
	}
	if petition.Category != "" {
		existingPetition.Category = petition.Category
	}
	if petition.Description != "" {
		existingPetition.Description = petition.Description
	}
	if petition.Image != "" {
		existingPetition.Image = petition.Image
	}
	if petition.VoteGoal != 0 {
		existingPetition.VoteGoal = petition.VoteGoal
	}
	if petition.ExpDate.IsZero() {
		existingPetition.ExpDate = petition.ExpDate
	}

	existingPetition.UpdatedAt = time.Now()
	if err := repo.db.Save(existingPetition).Error; err != nil {
		return err
	}

	return nil
}

func (repo *PetitionRepository) Delete(uuid string) error {
	var petition models.Petition
	err := repo.db.Unscoped().Where("uuid = ?", uuid).Delete(&petition).Error
	if err != nil {
		return err
	} else if petition.UUID == "0" {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *PetitionRepository) GetStatusByTitle(title string) (models.Status, error) {
	var status models.Status
	err := repo.db.Where("title = ?", title).First(&status).Error
	if err != nil {
		return models.Status{}, err
	}
	return status, nil
}

func (repo *PetitionRepository) GetByID(uuid string) (models.Petition, error) {
	var petition models.Petition
	err := repo.db.Where("uuid = ?", uuid).Preload("Status").First(&petition).Error
	if err != nil {
		return petition, err
	}
	return petition, nil
}

func (repo *PetitionRepository) CheckIfExists(uuid string) error {
	var petitions models.Petition
	if err := repo.db.Where("uuid = ?", uuid).First(&petitions).Error; err != nil {
		return err
	}

	return nil
}

func (repo *PetitionRepository) HasUserVoted(userID, petitionID string) error {
	var vote models.Vote
	if err := repo.db.Where("petition_id = ?", petitionID).First(&vote).Error; err != nil {
		return nil
	}

	if vote.UserID == userID {
		return nil
	}

	return errors.New("user has already voted")
}

func (repo *PetitionRepository) GetAllUserPetitions(userID string, pagination util.Pagination) ([]models.Petition, error) {
	var petitions []models.Petition
	if err := repo.db.Scopes(postgres.Paginate(pagination)).Model(models.Petition{}).Where("user_id = ?", userID).Preload("Status").Find(&petitions).Error; err != nil {
		return nil, err
	}
	return petitions, nil
}

func (repo *PetitionRepository) GetAllUserVotedPetitions(userID string, pagination util.Pagination) ([]models.Petition, error) {
	var petitions []models.Petition

	query := `
        SELECT petitions.*
        FROM petitions
        JOIN votes ON petitions.uuid = votes.petition_uuid
        WHERE votes.user_id = ?
        LIMIT ? OFFSET ?
    `
	if err := repo.db.Raw(query, userID, pagination.Limit, pagination.Page).Scan(&petitions).Error; err != nil {
		return nil, err
	}

	return petitions, nil
}

func (r *PetitionRepository) UpdateCurrVotes(petition models.Petition) error {
	if err := r.db.Save(&petition).Error; err != nil {
		return err
	}
	return nil
}

func (repo *PetitionRepository) GetPetitionsTitles(pagination util.Pagination) ([]models.PetitionInfo, error) {
	var petitionInfo []models.PetitionInfo

	err := repo.db.Debug().Scopes(postgres.Paginate(pagination)).Table("petitions").Select("id, user_id, title,  description,  author_name").Find(&petitionInfo).Error
	if err != nil {
		return nil, err
	}

	return petitionInfo, nil
}

func (repo *PetitionRepository) SearchPetitionsByTitle(searchTerm string, pagination util.Pagination) ([]models.PetitionInfo, error) {
	var petitions []models.PetitionInfo
	searchTerm = "%" + searchTerm + "%"
	err := repo.db.Where("lower(title) LIKE lower(?)", searchTerm).Table("petitions").Scopes(postgres.Paginate(pagination)).
		Select("id, user_id, title,  description,  author_name").Find(&petitions).Error
	if err != nil {
		return nil, err
	}
	return petitions, nil
}
