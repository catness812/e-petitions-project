package models

type User struct {
	Id       int    `json:"id" gorm:"primarykey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"Role"`
}
