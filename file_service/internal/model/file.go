package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	UserID   string
	FileType string
	Chunks   []Chunk
}

type User struct {
	Id         uint32 `gorm:"primaryKey;autoIncrement:true;unique"`
	Email      string `gorm:"not null;unique" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	Role       string `json:"role"`
	HasAccount bool   `json:"HasAccount"`
	Files      []File `gorm:"many2many:user_files"`
}

type Chunk struct {
	ID             uint   `gorm:"primaryKey"`
	SequenceNumber int    `gorm:"not null"`
	Data           []byte `gorm:"not null"`
	FileID         uint   `gorm:"not null"`
}
