package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	UserID   string
	FileType string
	Chunks   []Chunk
}

type Chunk struct {
	ID             uint   `gorm:"primaryKey"`
	SequenceNumber int    `gorm:"not null"`
	Data           []byte `gorm:"not null"`
	FileID         uint   `gorm:"not null"`
}
