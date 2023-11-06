package model

type User struct {
	UUID string `json:"id"`
	Role string `json:"role"`
	UserCredentials
}

type UserCredentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserUUID     string `json:"user_id"`
}
