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

type CreatePetition struct {
	// string title = 1;
	// string description = 2;
	// string image = 3;
	// uint32 user_id = 4;
	// string category = 5;
	// uint32 vote_goal = 8;
	// Title string `json:"title"`
}
