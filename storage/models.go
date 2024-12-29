package storage

type Word struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Lang         string `gorm:"size:5;index:idx_lang_chars"`
	OriginalWord string `gorm:"size:20"`
	LatinWord    string `gorm:"size:10;index"`
	Chars        byte   `gorm:"index:idx_lang_chars"`
}
