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
	r.DB.AutoMigrate(&Collection{}, &Word{})
}

func (r *Repository) PurgeDatabase() {
	os.Remove(r.DBName)
}

func (r *Repository) FindCollectionByCode(collection *Collection, code string) error {
	result := r.DB.First(collection, "code = ?", code)
	return result.Error
}

func (r *Repository) LatinWordExists(collection *Collection, latinWord string) bool {
	foundWord := &Word{}
	r.DB.First(foundWord, "collection_id = ? AND latin_word = ?", &collection.ID, latinWord)

	return foundWord.ID > 0
}

func (r *Repository) CreateWord(word *Word) error {
	result := r.DB.Create(&word)
	return result.Error
}

func (r *Repository) CreateCollection(collection *Collection) error {
	result := r.DB.Create(&collection)
	return result.Error
}

func (r *Repository) FindRandomWords(collection *Collection, chars int, limit int) (*[]Word, error) {
	words := &[]Word{}
	result := r.DB.Limit(limit).Order("RANDOM()").
		Find(words, "collection_id = ? AND chars = ?", &collection.ID, chars)

	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}
