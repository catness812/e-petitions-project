package models

type User struct {
	UUID       string `gorm:"primaryKey;autoIncrement:true;unique"`
	Email      string `gorm:"not null;unique" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	Role       string `json:"role"`
	HasAccount bool   `json:"HasAccount"`
}

type UserCredentialsModel struct {
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password"`
}
