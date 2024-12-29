package storage

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	DBName string
	DB     *gorm.DB
}

func (r *Repository) OpenDatabase() {
	db, err := gorm.Open(sqlite.Open(r.DBName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	r.DB = db
}

func (r *Repository) PrepareDatabase() {
	r.OpenDatabase()
	r.DB.AutoMigrate(&Word{})
}

func (r *Repository) PurgeDatabase() {
	os.Remove(r.DBName)
}

func (r *Repository) LatinWordExists(lang string, latinWord string) bool {
	foundWord := &Word{}
	r.DB.First(foundWord, "latin_word = ? AND lang = ?", latinWord, lang)

	return foundWord.ID > 0
}

func (r *Repository) CreateWord(word *Word) error {
	result := r.DB.Create(&word)
	return result.Error
}

func (r *Repository) FindRandomWords(lang string, chars int, limit int) (*[]Word, error) {
	words := &[]Word{}
	result := r.DB.Limit(limit).Order("RANDOM()").
		Find(words, "lang = ? AND chars = ?", lang, chars)

	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}
