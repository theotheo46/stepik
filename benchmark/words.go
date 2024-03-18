package main

// не удаляйте импорты, они используются при проверке
import (
	"strings"
)

type Words struct {
	str   string
	words []string
}

func MakeWords(s string) Words {
	words := strings.Fields(s)
	return Words{s, words}
}

func (w Words) Index(word string) int {
	for idx, item := range w.words {
		if item == word {
			return idx
		}
	}
	return -1
}
