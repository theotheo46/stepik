package main

// не удаляйте импорты, они используются при проверке
import (
	"strings"
	"testing"
)

// используйте эту переменную в бенчмарках
const src = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

// используйте эту переменную в бенчмарках
const pattern = "commodo"

// реализуйте бенчмарк для MatchContains
func BenchmarkMatchContains(b *testing.B) {
	// ...
	for n := 0; n < b.N; n++ {
		MatchContains(pattern, src)
	}
}

// реализуйте бенчмарк для MatchContainsCustom
func BenchmarkMatchContainsCustom(b *testing.B) {
	// ...
	for n := 0; n < b.N; n++ {
		MatchContainsCustom(pattern, src)
	}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// Библиотечная
func MatchContains(pattern string, src string) bool {
	return strings.Contains(src, pattern)
}

// Самописная
func MatchContainsCustom(pattern string, src string) bool {
	if pattern == "" {
		return true
	}
	if len(pattern) > len(src) {
		return false
	}
	pat_len := len(pattern)
	for idx := 0; idx < len(src)-pat_len+1; idx++ {
		if src[idx:idx+pat_len] == pattern {
			return true
		}
	}
	return false
}
