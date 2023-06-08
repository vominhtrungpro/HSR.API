package charactermodel

type SearchResult struct {
	ID      int32  `gorm:"column:id;primaryKey" json:"id"`
	Name    string `gorm:"column:name;not null" json:"name"`
	Rarity  int32  `gorm:"column:rarity;not null" json:"rarity"`
	Element int32  `gorm:"column:element;not null" json:"element"`
	Path    int32  `gorm:"column:path;not null" json:"path"`
	Url     string `gorm:"column:path;not null" json:"url"`
}
