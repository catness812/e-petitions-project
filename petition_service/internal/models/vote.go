package models

type Vote struct {
	ID         uint `gorm:"primaryKey;autoIncrement:true"`
	UserID     uint `gorm:"not null"`
	PetitionID uint `gorm:"not null"`
}
