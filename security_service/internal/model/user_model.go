package models

type UserModel struct {
	ID       uint32 `gorm:"primaryKey;autoIncrement:true;unique"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (UserModel) TableName() string {
	return "users"
}

type UserCredentialsModel struct {
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password" `
}
