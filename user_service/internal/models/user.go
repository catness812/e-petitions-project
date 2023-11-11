package models

type User struct {
	ID         uint   `gorm:"primaryKey;autoIncrement:true;unique"`
	UUID       string `gorm:"unique"`
	Email      string `gorm:"not null;unique" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	Role       string `json:"role"`
	HasAccount bool   `json:"HasAccount"`
}
