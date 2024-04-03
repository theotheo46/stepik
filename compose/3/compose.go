package main

import (
	"fmt"
	"math/rand"
)

// начало решения

// генерит случайные слова из 5 букв
// с помощью randomWord(5)

func ReverseString(s string) string {
	runes := []rune(s)
	size := len(runes)
	for i, j := 0, size-1; i < size>>1; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isCountMoreOne(str string) bool {
	frequency := make(map[rune]int)
	for _, char := range str {
		frequency[char] = frequency[char] + 1
	}
	for _, val := range frequency {
		if val > 1 {
			return true
		}
	}
	return false
}

func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case out <- randomWord(5):
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for word := range in {
			if isCountMoreOne(word) {
				out <- word
			}
		}
	}()
	return out
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case word, _ := <-in:
				out <- ReverseString(word)
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, in1, in2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for in1 != nil || in2 != nil {
			select {
			case val1, ok := <-in1:
				if ok {
					//fmt.Println(val1)
					out <- val1
				} else {
					in1 = nil
				}
			case val2, ok := <-in2:
				if ok {
					out <- val2
					//fmt.Println(val2)
				} else {
					in2 = nil
				}
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(<-in)
	}
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
