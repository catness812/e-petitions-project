package models

type User struct {
	Id       int    `gorm:"primaryKey;autoIncrement:true;unique"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     string `json:"Role"`
}
