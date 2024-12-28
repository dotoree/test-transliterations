package storage

type Word struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	OriginalWord string `gorm:"size:20"`
	LatinWord    string `gorm:"size:10;index"`
	Chars        byte   `gorm:"index"`
}
