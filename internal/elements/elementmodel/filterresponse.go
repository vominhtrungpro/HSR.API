package elementmodel

type FilterResponse struct {
	Name   string `gorm:"column:name;not null" json:"name"`
	Enname string `gorm:"column:enname;not null" json:"enname"`
	Type   string `gorm:"column:type;not null" json:"type"`
}
