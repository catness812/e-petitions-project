package model

type File struct {
	Id         uint32 `gorm:"primaryKey;autoIncrement:true;unique"`
	UserID     uint32 `gorm:"not null"`
	PetitionID uint32
	Type       string `gorm:"not null"`
	Data       []byte `gorm:"not null"`
}

type User struct {
	Id         uint32 `gorm:"primaryKey;autoIncrement:true;unique"`
	Email      string `gorm:"not null;unique" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	Role       string `json:"role"`
	HasAccount bool   `json:"HasAccount"`
}

type UserPhoto struct {
	Id     uint32 `gorm:"primaryKey;autoIncrement:true;unique"`
	UserID uint32
	User   User `gorm:"foreignKey:UserID"`
	FileID uint32
	File   File `gorm:"foreignKey:FileID"`
}
