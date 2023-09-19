package models

import "gorm.io/gorm"

type Petition struct {
	gorm.Model
	Title       string `gorm:"not null;" json:"title"`
	Category    string `gorm:"not null;" json:"category"`
	Description string `gorm:"not null;" json:"description"`
	Image       string `gorm:"not null;" json:"image"`
	// we could use base64 encoding
}
