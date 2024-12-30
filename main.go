package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strconv"
	"strings"

	"github.com/dotoree/test-transliterations/storage"
	"github.com/dotoree/test-transliterations/utils"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

var repo *storage.Repository

var dictionaries = []utils.Dictionary{
	{
		Lang:     "el",
		Filename: "Greek_utf8.dic",
	},
	{
		Lang:     "en",
		Filename: "English.dic",
	},
}

func main() {
	repo = &storage.Repository{
		DBName: "bybonpass.db",
		DB:     &gorm.DB{},
	}

	// Imports
	// importLangDictionary("el")
	// importLangDictionary("en")
	importLangDictionary("xxx")

	// Test results
	repo.OpenDatabase()
	lang := "en"
	words, err := repo.FindRandomWords(lang, 6, 4)

	if err != nil {
		panic(err)
	}

	p := []string{}
	for _, word := range *words {
		fmt.Printf("%#v\n", word)
		w := word.LatinWord
		p = append(p, w)
	}

	// Add number
	p = append(p, strconv.Itoa(rand.IntN(1000)+1))

	// Add caser for title case first word
	caser := cases.Title(language.English)
	p[0] = caser.String(p[0])

	// Join
	password := strings.Join(p[:], "-")

	// Validate
	entropy := passwordvalidator.GetEntropy(password)

	fmt.Printf("\n\nPassword: %s\n", password)
	fmt.Printf("Entropy: %v\n\n", entropy)
}

func importLangDictionary(lang string) {
	idx := slices.IndexFunc(dictionaries, func(d utils.Dictionary) bool { return d.Lang == lang })

	if idx > -1 {
		utils.ImportDictionary(&dictionaries[idx], repo, 0)
	}
}
