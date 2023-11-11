package repository

import (
	"github.com/catness812/e-petitions-project/file_service/internal/model"
	"gorm.io/gorm"
)

type FileRepository struct {
	dbClient *gorm.DB
}

func NewFileRepository(dbClient *gorm.DB) *FileRepository {
	return &FileRepository{dbClient: dbClient}
}

func (repo *FileRepository) StoreFile(file *model.File) (uint32, error) {
	if err := repo.dbClient.Debug().Model(model.File{}).Create(file).Error; err != nil {
		return 0, err
	}
	return file.Id, nil
}

func (repo *FileRepository) StoreUserPic(pic *model.UserPhoto) error {
	if err := repo.dbClient.Debug().Model(model.UserPhoto{}).Create(pic).Error; err != nil {
		return err
	}
	return nil
}

func (repo *FileRepository) FetchUserPic(userID uint32) (*model.File, error) {
	pic := &model.UserPhoto{}
	if err := repo.dbClient.Debug().Model(model.UserPhoto{}).Where("user_id = ?", userID).First(pic).Error; err != nil {
		return nil, err
	}
	file := &model.File{}
	if err := repo.dbClient.Debug().Model(model.File{}).Where("id = ?", pic.FileID).First(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func (repo *FileRepository) FetchFile(fileID uint32) (*model.File, error) {
	file := &model.File{}
	if err := repo.dbClient.Debug().Model(model.File{}).Where("id = ?", fileID).First(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func (repo *FileRepository) FetchPetitionFiles(petitionID uint32) ([]model.File, error) {
	var files []model.File
	if err := repo.dbClient.Debug().Model(model.File{}).Where("petition_id = ?", petitionID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}
