package service

import "github.com/catness812/e-petitions-project/file_service/internal/model"

type IFileRepository interface {
	StoreFile(file *model.File) (uint32, error)
	StoreUserPic(pic *model.UserPhoto) error
	FetchUserPic(userID uint32) (*model.File, error)
	FetchFile(fileID uint32) (*model.File, error)
	FetchPetitionFiles(petitionID uint32) ([]model.File, error)
}

type FileService struct {
	repo IFileRepository
}

func NewFileService(repo IFileRepository) *FileService {
	return &FileService{repo: repo}
}

func (fileSvc *FileService) SendFile(file *model.File) error {
	if _, err := fileSvc.repo.StoreFile(file); err != nil {
		return err
	}
	return nil
}

func (fileSvc *FileService) SendUserPic(file *model.File) error {
	fileID, err := fileSvc.repo.StoreFile(file)
	if err != nil {
		return err
	}
	pic := &model.UserPhoto{
		UserID: file.UserID,
		FileID: fileID,
	}
	if err := fileSvc.repo.StoreUserPic(pic); err != nil {
		return err
	}
	return nil
}

func (fileSvc *FileService) GetUserPic(userID uint32) (*model.File, error) {
	pic, err := fileSvc.repo.FetchUserPic(userID)
	if err != nil {
		return nil, err
	}
	return pic, err
}

func (fileSvc *FileService) GetFile(fileID uint32) (*model.File, error) {
	file, err := fileSvc.repo.FetchFile(fileID)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (fileSvc *FileService) GetPetitionFiles(petitionID uint32) ([]model.File, error) {
	files, err := fileSvc.repo.FetchPetitionFiles(petitionID)
	if err != nil {
		return nil, err
	}
	return files, nil
}
