package main

import (
	"strings"
	"testing"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func isSafeSymbol(b rune) bool {
	switch {
	case b >= 97 && b <= 122:
		return true
	case b >= 65 && b <= 90:
		return true
	case b >= 48 && b <= 57:
		return true
	case b == 45:
		return true
	default:
		return false
	}
}

func slugify(src string) string {
	ret := strings.Clone(src)
	for _, s := range src {
		if !isSafeSymbol(s) {
			ret = strings.ReplaceAll(ret, string(s), " ")
		}
	}
	return strings.Join(strings.Fields(strings.ToLower(ret)), "-")
}

// конец решения

func Test(t *testing.T) {
	const phrase = "Go Is Awesome!"
	const want = "go-is-awesome"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}
