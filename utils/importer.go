package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"

	"strings"

	"github.com/mozillazg/go-unidecode"
	"github.com/stts-se/translit/grc"

	"github.com/dotoree/test-transliterations/storage"
)

type Dictionary struct {
	Lang     string
	Filename string
}

func ImportDictionary(d *Dictionary, r *storage.Repository, maxEntries int) {
	r.PrepareDatabase()

	// Open Greek dictionary file
	file, err := os.Open(d.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		greekWord := strings.ToLower(strings.TrimSpace(scanner.Text()))

		// Skip numbers
		if _, err := strconv.Atoi(greekWord); err == nil {
			continue
		}

		// Transliterate
		latinWord, err := grc.Convert(greekWord)
		if err != nil {
			log.Println("Could not transliterate word:", greekWord, err)
			continue
		}

		// Convert to unicode
		latinWord = unidecode.Unidecode(latinWord)

		// Skip non ASCII words
		if !isASCII(latinWord) {
			continue
		}

		chars := len(latinWord)
		if chars >= 4 && chars <= 8 {
			// Skip duplicate words
			if r.LatinWordExists(latinWord) {
				continue
			}

			word := storage.Word{OriginalWord: greekWord, LatinWord: latinWord, Chars: byte(chars)}
			if err = r.CreateWord(&word); err != nil {
				log.Println("Error inserting word:", latinWord, err)
			}
			counter++
			if maxEntries > 0 && counter >= maxEntries {
				break
			}
		}
	}

	fmt.Printf("%d Words inserted successfully.", counter)
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
