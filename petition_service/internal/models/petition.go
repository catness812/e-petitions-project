package models

import (
	"time"

	"gorm.io/gorm"
)

type Petition struct {
	gorm.Model
	Title       string    `gorm:"not null;" json:"title"`
	Category    string    `gorm:"not null;" json:"category"`
	Description string    `gorm:"not null;" json:"description"`
	Image       string    `gorm:"not null;" json:"image"`
	StatusID    uint      `gorm:"not null;" json:"status_id"`
	Status      Status    `gorm:"foreignKey:StatusID" json:"status"`
	UserID      uint      `gorm:"not null;" json:"user_id"`
	AuthorName  string    `gorm:"not null;" json:"author_name"`
	VoteGoal    uint      `gorm:"not null;default:1000" json:"vote_goal"`
	CurrVotes   uint      `gorm:"not null;default:0" json:"curr_votes"`
	ExpDate     time.Time `gorm:"not null;" json:"exp_date"`
}

type PetitionInfo struct {
	gorm.Model
	Title       string `gorm:"not null;" json:"title"`
	Description string `gorm:"not null;" json:"description"`
	UserID      uint   `gorm:"not null;" json:"user_id"`
	AuthorName  string `gorm:"not null;" json:"author_name"`
	ID          uint   `gorm:"not null;" json:"id"`
}

type PetitionUpdate struct {
	ID          uint      `gorm:"not null;" json:"id"`
	Title       string    `gorm:"not null;" json:"title"`
	Category    string    `gorm:"not null;" json:"category"`
	Description string    `gorm:"not null;" json:"description"`
	Image       string    `gorm:"not null;" json:"image"`
	StatusID    uint      `gorm:"not null;" json:"status_id"`
	VoteGoal    uint      `gorm:"not null;default:1000" json:"vote_goal"`
	ExpDate     time.Time `gorm:"not null;" json:"exp_date"`
}
