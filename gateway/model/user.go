package model

type User struct {
	Id   uint32 `json:"id"`
	Role string `json:"role"`
	UserCredentials
}

type UserCredentials struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8,max=50"`
	HasAccount bool   `json:"hasAccount"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
