package repository

import "gorm.io/gorm"

type FileRepository struct {
	dbClient *gorm.DB
}

func NewFileRepository(dbClient *gorm.DB) *FileRepository {
	return &FileRepository{dbClient: dbClient}
}

func (fileRepo *FileRepository) ProcessAndStoreChunk(fileID, sequenceNumber int, chunk []byte) error {

}
