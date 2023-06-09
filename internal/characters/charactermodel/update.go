package charactermodel

type UpdateRequest struct {
	Id      string `gorm:"column:id;not null" json:"id"`
	Name    string `gorm:"column:name;not null" json:"name"`
	Rarity  int32  `gorm:"column:rarity;not null" json:"rarity"`
	Element int32  `gorm:"column:element;not null" json:"element"`
	Path    int32  `gorm:"column:path;not null" json:"path"`
}
