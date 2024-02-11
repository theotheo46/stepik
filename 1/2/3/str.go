package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	result := ""
	for i := 0; i < times; i++ {
		result += source
	}

	fmt.Println(result)
}
