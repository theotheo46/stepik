package main

import (
	"fmt"
	"os"
	"strings"
)

// начало решения

// readLines возвращает все строки из указанного файла
func readLines(name string) ([]string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("file name %s is not found", name)
	}

	sarr_splitted := strings.Split(string(data[:]), "\n")
	if sarr_splitted[len(sarr_splitted)-1] == "" {
		sarr_splitted = sarr_splitted[:len(sarr_splitted)-1]
	}
	return sarr_splitted, nil
}

// конец решения

func main() {
	lines, err := readLines("/etc/passwd")
	if err != nil {
		panic(err)
	}
	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx+1, line)
	}
}
