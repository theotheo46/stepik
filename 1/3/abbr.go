package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	phrase := readString()
	phrase_arr := strings.Fields(phrase)
	abbr := ""
	for _, str := range phrase_arr {
		r, _ := utf8.DecodeRuneInString(str)
		if unicode.IsLetter(r) {
			r = unicode.ToUpper(r)
			abbr += string(r)
		}
	}

	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
