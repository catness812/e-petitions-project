package model

type Petition struct {
	PetitionId    uint32 `gorm:"not null;" json:"petition_id"`
	Title         string `gorm:"not null;" json:"title"`
	Category      string `gorm:"not null;" json:"category"`
	Description   string `gorm:"not null;" json:"description"`
	Image         string `gorm:"not null;" json:"image"`
	Status        Status `gorm:"not null;" json:"status"`
	UserID        uint32 `gorm:"not null;" json:"user-id"`
	Vote_Goal     uint32 `gorm:"not null;" json:"vote_goal"`
	Current_Votes uint32 `gorm:"not null;" json:"current_votes"`
	Exp_Date      string `gorm:"not null;" json:"exp_date"`
	UpdatedAt     string `gorm:"not null;" json:"updated_at"`
	CreatedAt     string `gorm:"not null;" json:"created_at"`
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
	ID     uint32 `json:"id"`
	Status string `json:"status"`
}
