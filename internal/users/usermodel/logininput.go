package usermodel

type LoginInput struct {
	Username string `gorm:"column:username;not null" json:"username"`
	Password string `gorm:"column:password;not null" json:"password"`
}
