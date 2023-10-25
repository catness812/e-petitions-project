package repository

import (
	"errors"
	"time"

	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
	"github.com/catness812/e-petitions-project/petition_service/pkg/database/postgres"
	"github.com/gookit/slog"
	"gorm.io/gorm"
)

type PetitionRepository struct {
	db *gorm.DB
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

func NewPetitionRepository(db *gorm.DB) *PetitionRepository {
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

func (repo *PetitionRepository) SaveVote(Vote *models.Vote) error {
	err := repo.db.Create(Vote).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PetitionRepository) UpdateStatus(id uint, statusID uint) error {
	var petition models.Petition
	// first query to see if this petition exists
	result := repo.db.Where("id = ?", id).First(&petition)
	if result.Error != nil {
		return result.Error
	}
	petition.StatusID = statusID
	petition.UpdatedAt = time.Now()

	repo.db.Save(&petition)
	return nil
}

func (repo *PetitionRepository) Delete(id uint) error {
	var petition models.Petition
	err := repo.db.Unscoped().Where("id = ?", id).Delete(&petition).Error
	if err != nil {
		return err
	} else if petition.ID == 0 {
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

func (repo *PetitionRepository) GetByID(id uint) (models.Petition, error) {
	var petition models.Petition
	err := repo.db.Where("id = ?", id).Preload("Status").First(&petition).Error
	if err != nil {
		return petition, err
	}
	return petition, nil
}

func (repo *PetitionRepository) CheckIfExists(id uint) error {
	var petitions models.Petition
	if err := repo.db.Where("id = ?", id).First(&petitions).Error; err != nil {
		slog.Errorf("Couldn't find petition: %v", err.Error())
		return err
	}

	slog.Infof("petition found")
	return nil
}

func (repo *PetitionRepository) HasUserVoted(userID, petitionID uint) error {
	var vote models.Vote
	if err := repo.db.Where("user_id = ? AND petition_id = ?", userID, petitionID).First(&vote).Error; err != nil {
		slog.Info("Couldn't find vote")
		return nil
	}
	slog.Error("Vote found")
	return errors.New("user has already voted")
}
func (repo *PetitionRepository) GetAllUserPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error) {
	var petitions []models.Petition
	if err := repo.db.Scopes(postgres.Paginate(pagination)).Model(models.Petition{}).Where("user_id = ?", userID).Preload("Status").Find(&petitions).Error; err != nil {
		return nil, err
	}
	return petitions, nil
}

func (repo *PetitionRepository) GetAllUserVotedPetitions(userID uint, pagination util.Pagination) ([]models.Petition, error) {
	var petitions []models.Petition

	query := `
        SELECT petitions.*
        FROM petitions
        JOIN votes ON petitions.id = votes.petition_id
        WHERE votes.user_id = ?
        LIMIT ? OFFSET ?
    `
	if err := repo.db.Raw(query, userID, pagination.Limit, pagination.Page).Scan(&petitions).Error; err != nil {
		slog.Errorf("can't execute the query: %v", err)
		return nil, err
	}
	slog.Info(petitions)

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

	err := repo.db.Debug().Scopes(postgres.Paginate(pagination)).Table("petitions").Select("id, user_id, title").Find(&petitionInfo).Error
	if err != nil {
		return nil, err
	}

	return petitionInfo, nil
}

func (repo *PetitionRepository) SearchPetitionsByTitle(searchTerm string, pagination util.Pagination) ([]models.PetitionInfo, error) {
	var petitions []models.PetitionInfo
	searchTerm = "%" + searchTerm + "%"
	err := repo.db.Where("lower(title) LIKE lower(?)", searchTerm).Table("petitions").Scopes(postgres.Paginate(pagination)).
		Select("id, user_id, title").Find(&petitions).Error
	if err != nil {
		return nil, err
	}
	return petitions, nil
}
