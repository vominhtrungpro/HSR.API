package usermodel

type LoginOutput struct {
	AccessToken  string `gorm:"column:accesstoken;not null" json:"accesstoken"`
	RefreshToken string `gorm:"column:refreshtoken;not null" json:"refreshtoken"`
}
