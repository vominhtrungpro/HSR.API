package usermodel

// User mapped from table <users>
type CreateUserRequest struct {
	Username string `gorm:"column:username;not null" json:"username"`
	Password string `gorm:"column:password;not null" json:"password"`
	Email    string `gorm:"column:email;not null" json:"email"`
}
