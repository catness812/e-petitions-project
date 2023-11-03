package models

import (
	"time"

	"gorm.io/gorm"
)

type Petition struct {
	gorm.Model
	UUID        string    `gorm:"unique"`
	Title       string    `gorm:"not null;" json:"title"`
	Category    string    `gorm:"not null;" json:"category"`
	Description string    `gorm:"not null;" json:"description"`
	Image       string    `gorm:"not null;" json:"image"`
	StatusID    uint      `gorm:"not null;" json:"status_id"`
	Status      Status    `gorm:"foreignKey:StatusID" json:"status"`
	UserID      string    `gorm:"not null;" json:"user_id"`
	AuthorName  string    `gorm:"not null;" json:"author_name"`
	VoteGoal    uint      `gorm:"not null;default:1000" json:"vote_goal"`
	CurrVotes   uint      `gorm:"not null;default:0" json:"curr_votes"`
	ExpDate     time.Time `gorm:"not null;" json:"exp_date"`
	Votes       []Vote    `gorm:"foreignKey:PetitionUUID" json:"votes"`
}

type Vote struct {
	ID           uint     `gorm:"primaryKey;autoIncrement:true"`
	UserID       string   `gorm:"not null" json:"user_id"`
	PetitionUUID string   `gorm:"not null" json:"petition_uuid"`
	Petition     Petition `gorm:"foreignKey:PetitionUUID" json:"petition"`
}

type PetitionInfo struct {
	gorm.Model
	Title       string `gorm:"not null;" json:"title"`
	Description string `gorm:"not null;" json:"description"`
	UserID      string `gorm:"not null;" json:"user_id"`
	AuthorName  string `gorm:"not null;" json:"author_name"`
	UUID        string `gorm:"not null;"`
}

type PetitionUpdate struct {
	UUID        string    `gorm:"not null;"`
	Title       string    `gorm:"not null;" json:"title"`
	Category    string    `gorm:"not null;" json:"category"`
	Description string    `gorm:"not null;" json:"description"`
	Image       string    `gorm:"not null;" json:"image"`
	StatusID    uint      `gorm:"not null;" json:"status_id"`
	VoteGoal    uint      `gorm:"not null;default:1000" json:"vote_goal"`
	ExpDate     time.Time `gorm:"not null;" json:"exp_date"`
}
