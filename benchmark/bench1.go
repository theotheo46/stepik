package main

// не удаляйте импорты, они используются при проверке
import (
    "fmt"
    "math/rand"
    "os"
    "testing"
)

// не удаляйте импорты, они используются при проверке

// реализуйте быстрое множество
type IntSet struct {
	elems *map[int]any
}

func MakeIntSet() IntSet {
	elems := make(map[int]any)
	return IntSet{&elems}
}

func (s IntSet) Contains(elem int) bool {
	_, ok := (*s.elems)[elem]
	if ok == true {
		return true
	} else {
		return false
	}
}

func (s IntSet) Add(elem int) bool {
	if s.Contains(elem) {
		return false
	}
	(*s.elems)[elem] = nil
	return true
}
