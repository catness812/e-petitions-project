package model

type Petition struct {
	PetitionId  uint32 `gorm:"not null;" json:"petition_id"`
	Title       string `gorm:"not null;" json:"title"`
	Category    string `gorm:"not null;" json:"category"`
	Description string `gorm:"not null;" json:"description"`
	Image       string `gorm:"not null;" json:"image"`
	Status      Status `gorm:"not null;" json:"status"`
	UserID      uint32 `gorm:"not null;" json:"user-id"`
}

type CreatePetition struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      uint32 `json:"user_id"`
	Category    string `json:"category"`
	VoteGoal    uint32 `json:"vote_goal"`
}

type Status struct {
	ID    uint32 `json:"id"`
	Title string `json:"titile"`
}
