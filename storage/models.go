package storage

type Word struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	OriginalWord string `gorm:"size:20"`
	LatinWord    string `gorm:"size:10;index"`
	Chars        byte   `gorm:"index:idx_collection_chars"`
	CollectionID uint   `gorm:"index:idx_collection_chars"`
	Collection   Collection
}

type Collection struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Lang string `gorm:"size:5;index"`
	Code string `gorm:"size:20:index:unique"`
	Name string
}
