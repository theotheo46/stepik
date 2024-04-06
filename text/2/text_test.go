package main

import (
	"strconv"
	"strings"
	"testing"
)

// начало решения

// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
	total := 0
	for _, s := range directions {
		sarr := strings.Fields(s)
		for _, s1 := range sarr {
			if strings.HasSuffix(s1, "m") || strings.HasSuffix(s1, "km") {
				if strings.HasSuffix(s1, "m") && !strings.HasSuffix(s1, "km") {
					n, err := strconv.Atoi(strings.TrimRight(s1, "m"))
					if err == nil {
						total += n
					}
				} else {
					n, err := strconv.ParseFloat(strings.TrimRight(s1, "km"), 64)
					if err == nil {
						total += int(n * 1000)
					}
				}
			}
		}

	}
	return total
}

// конец решения

func Test(t *testing.T) {
	directions := []string{
		"100m to intersection",
		"turn right",
		"straight 300m",
		"enter motorway",
		"straight 5.1236km",
		"exit motorway",
		"500m straight",
		"turn sharp left",
		"continue 100m to destination",
	}
	const want = 6000
	got := calcDistance(directions)
	if got != want {
		t.Errorf("%v: got %v, want %v", directions, got, want)
	}
}
