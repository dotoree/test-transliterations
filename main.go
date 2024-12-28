package main

import (
	"fmt"
	"strings"

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
	words, err := repo.FindRandomWords(6, 4)

	if err != nil {
		panic(err)
	}

	fmt.Println("")
	p := []string{}
	for _, word := range *words {
		fmt.Printf("%#v\n", word)
		p = append(p, word.LatinWord)
	}
	fmt.Println("\n" + strings.Join(p[:], "-"))
	fmt.Println("")
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
