package elementmodel

type CreateRequest struct {
	Name   string `gorm:"column:name;not null" json:"name"`
	Enname string `gorm:"column:enname;not null" json:"enname"`
}
