package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

// начало решения

// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {
	//fmt.Println(t.UnixNano())
	//var d float64
	var ret string
	i := strconv.Itoa(int(t.UnixNano()))
	//fmt.Println("i: ", i)
	if i != "0" {
		d_part := i[len(i)-9:]
		if d_part == "000000000" {
			d_part = "0"
		} else {
			d_part = strings.TrimRight(d_part, "0")
		}
		i_part := i[0 : len(i)-9]
		//fmt.Println("i_part:", i_part, " d_part:", d_part)
		ret = i_part + "." + d_part
	} else {
		ret = "0.0"
	}
	/* 	fmt.Println("t.UnixNano()=", t.UnixNano())
	   	d := float64(t.UnixNano()) / 1000000000.0
	   	fmt.Println(d)
	   	s := strconv.FormatFloat(d, 'f', -1, 64)
	   	//fmt.Println("****** ", s)
	   	if !strings.Contains(s, ".") {
	   		s = s + ".0"
	   	} */
	//fmt.Println(d)
	return ret
}

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {
	re := regexp.MustCompile(`(^\d+)\.(\d+)`)
	groups := re.FindStringSubmatch(d)
	if len(groups) != 3 {
		return time.Time{}, fmt.Errorf("string %s is not parsable as date", d)
	}
	/* 	for _, group := range groups {
		fmt.Println(group)
	} */
	i_part := groups[1]
	d_part := groups[2]
	i1, _ := strconv.Atoi(i_part)
	i2, _ := strconv.Atoi(d_part)

	switch len(d_part) {
	case 3:
		i2 *= 1000000
	case 6:
		i2 *= 1000
	}
	return time.Unix(int64(i1), int64(i2)), nil
}

// конец решения

func Test_asLegacyDate(t *testing.T) {
	samples := map[time.Time]string{
		time.Date(1970, 1, 1, 1, 0, 0, 951205000, time.UTC): "3600.951205999",
		time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC):         "3600.0",
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC):         "0.0",
	}
	for src, want := range samples {
		got := asLegacyDate(src)
		if got != want {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}

func Test_parseLegacyDate(t *testing.T) {
	samples := map[string]time.Time{
		"3600.123":    time.Date(1970, 1, 1, 1, 0, 0, 123000000, time.UTC),
		"3600.0":      time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
		"0.0":         time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		"1.123456789": time.Date(1970, 1, 1, 0, 0, 1, 123456789, time.UTC),
	}
	for src, want := range samples {
		got, err := parseLegacyDate(src)
		if err != nil {
			t.Fatalf("%v: unexpected error", src)
		}
		if !got.Equal(want) {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}
