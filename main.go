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
		DBName: "bybonpass.db",
		DB:     &gorm.DB{},
	}

	// Imports (Disabled)
	if false {
		importGreek()
	}

	// Test results
	repo.OpenDatabase()
	lang := "el"
	words, err := repo.FindRandomWords(lang, 6, 4)

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

func importGreek() {
	// Import Greek dictionary
	utils.ImportDictionary(&utils.Dictionary{
		Lang:     "el",
		Filename: "Greek_utf8.dic",
	}, repo, 0)
}
