package models

type Vote struct {
	ID         uint `gorm:"primaryKey;autoIncrement:true"`
	UserID     uint `gorm:"not null" json:"user_id"`
	PetitionID uint `gorm:"not null" json:"petition_id"`
}
