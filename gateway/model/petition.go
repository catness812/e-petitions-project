package model

type Petition struct {
	PetitionId   uint32 `gorm:"not null;" json:"petition_id"`
	Title        string `gorm:"not null;" json:"title"`
	Category     string `gorm:"not null;" json:"category"`
	Description  string `gorm:"not null;" json:"description"`
	Image        string `gorm:"not null;" json:"image"`
	Status       Status `gorm:"not null;" json:"status"`
	UserID       uint32 `gorm:"not null;" json:"user_id"`
	VoteGoal     uint32 `gorm:"not null;" json:"vote_goal"`
	CurrentVotes uint32 `gorm:"not null;" json:"current_votes"`
	ExpDate      string `gorm:"not null;" json:"exp_date"`
	UpdatedAt    string `gorm:"not null;" json:"updated_at"`
	CreatedAt    string `gorm:"not null;" json:"created_at"`
}

type CreatePetition struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      uint32 `json:"user_id"`
	Category    string `json:"category"`
	VoteGoal    uint32 `json:"vote_goal"`
	ExpDate     string `json:"exp_date"`
}

type Status struct {
	ID     uint32 `json:"id"`
	Status string `json:"status"`
}
