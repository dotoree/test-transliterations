package main

import (
	"github.com/dotoree/test-transliterations/storage"
	"github.com/dotoree/test-transliterations/utils"
	"gorm.io/gorm"
)

var repo *storage.Repository

func main() {
	repo = &storage.Repository{
		DBName: "el.db",
		DB:     &gorm.DB{},
	}

	// Purge existing database
	repo.PurgeDatabase()

	// Import Greek dictionary
	utils.ImportDictionary(&utils.Dictionary{
		Lang:     "el",
		Filename: "Greek_utf8.dic",
	}, repo, 1000)
}
