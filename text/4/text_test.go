package main

import (
	"regexp"
	"strings"
	"testing"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func slugify(src string) string {
	re := regexp.MustCompile(`[a-zA-Z0-9-]+`)
	sss := re.FindAllString(strings.ToLower(src), -1)
	return strings.Join(sss, "-")
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
