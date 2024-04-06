package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// начало решения

// prettify возвращает отформатированное
// строковое представление карты
func prettify(m map[string]int) string {
	var b strings.Builder
	if len(m) <= 1 {
		b.WriteString("{")
		for k, v := range m {
			b.WriteString(" ")
			b.WriteString(k)
			b.WriteString(": ")
			b.WriteString(strconv.Itoa(v))
			b.WriteString(" ")
		}
		b.WriteString("}")
	} else {
		keys := make([]string, 0, len(m))

		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		b.WriteString("{\n")
		for _, k := range keys {
			b.WriteString("    ")
			b.WriteString(k)
			b.WriteString(": ")
			b.WriteString(strconv.Itoa(m[k]))
			b.WriteString(",\n")
		}
		b.WriteString("}")
	}
	return b.String()
}

// конец решения

func Test(t *testing.T) {
	//m := map[string]int{"one": 1, "two": 2, "three": 3}
	m := map[string]int{}
	const want = "{\n    one: 1,\n    three: 3,\n    two: 2,\n}"
	got := prettify(m)
	fmt.Print(want)
	if got != want {
		t.Errorf("%v\ngot:\n%v\n\nwant:\n%v", m, got, want)
	}
}
