package main

import (
	"bufio"
	"fmt"
	"os"
)

// начало решения

// readLines возвращает все строки из указанного файла
func readLines(name string) ([]string, error) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("file name %s is not found", name)
	}
	sarr := []string{}
	scanner := bufio.NewScanner(file) // (1)
	for scanner.Scan() {              // (2)
		sarr = append(sarr, scanner.Text()) // (3)
	}
	return sarr, nil
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
