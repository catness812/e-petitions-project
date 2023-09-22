package model

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type UserCredentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}
