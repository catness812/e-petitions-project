package model

type Petition struct {
	PetitionId  uint32 `gorm:"not null;" json:"petition_id"`
	Title       string `gorm:"not null;" json:"title"`
	Category    string `gorm:"not null;" json:"category"`
	Description string `gorm:"not null;" json:"description"`
	Image       string `gorm:"not null;" json:"image"`
	Status      uint32 `gorm:"not null;" json:"status"`
	UserID      uint   `gorm:"not null;" json:"user-id"`
}
