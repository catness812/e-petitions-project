package models

import "gorm.io/gorm"

// status types
const (
	DRAFT     = "DRAFT"
	IN_REVIEW = "IN_REVIEW"
	PUBLIC    = "PUBLIC"
	ARCHIVE   = "ARCHIVE"
)

type Status struct {
	gorm.Model
	Title string `gorm:"not null;unique" json:"title"`
}

var StatusSeedData = []Status{
	{Title: ARCHIVE},
	{Title: IN_REVIEW},
	{Title: PUBLIC},
	{Title: DRAFT},
}
