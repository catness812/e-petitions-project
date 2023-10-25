package repository

import (
	"github.com/catness812/e-petitions-project/file_service/internal/model"
	"github.com/gookit/slog"
	"gorm.io/gorm"
)

type FileRepository struct {
	dbClient *gorm.DB
}

func NewFileRepository(dbClient *gorm.DB) *FileRepository {
	return &FileRepository{dbClient: dbClient}
}

func (fileRepo *FileRepository) ProcessAndStoreChunk(fileID uint, sequenceNumber int, chunk []byte) error {
	chunkData := &model.Chunk{FileID: fileID, SequenceNumber: sequenceNumber, Data: chunk}
	if err := fileRepo.dbClient.Debug().Model(model.Chunk{}).Create(chunkData).Error; err != nil {
		slog.Errorf("failed to insert chunk: %v", err)
		return err
	}
	return nil
}

func (fileRepo *FileRepository) CreateFile() (uint, error) {
	file := &model.File{}
	if err := fileRepo.dbClient.Debug().Model(model.File{}).Create(file).Error; err != nil {
		slog.Errorf("failed to insert file: %v", err)
		return 0, err
	}
	return file.ID, nil
}
