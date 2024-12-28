package main

import (
	"fmt"

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

	// Imports (Disabled)
	if false {
		doImports()
	}

	// Test results
	repo.OpenDatabase()
	word, err := repo.FindRandomWord(6)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", word)
}

func doImports() {
	// Purge existing database
	repo.PurgeDatabase()

	// Import Greek dictionary
	utils.ImportDictionary(&utils.Dictionary{
		Lang:     "el",
		Filename: "Greek_utf8.dic",
	}, repo, 0)
}
