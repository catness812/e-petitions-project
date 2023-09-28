package models

type User struct {
	Id       uint32 `gorm:"primaryKey;autoIncrement:true;unique"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     string `json:"Role"`
}
