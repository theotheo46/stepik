// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	count, _ := readInput()
	fmt.Println(count)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (size int, err error) {
	flag.Parse()
	src := strings.Join(flag.Args(), "")
	if src == "" {
		return 0, nil
	}
	return len(strings.Split(src, " ")), nil
}
