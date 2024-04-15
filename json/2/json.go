package main

import (
	"encoding/json"
	"fmt"
)

// начало решения

// Genre описывает жанр фильма
type Genre string

func IndexOfSubstring(str, subStr string) int {
	for i := 0; i < len(str); i++ {
		if str[i:i+len(subStr)] == subStr {
			return i
		}
	}
	return -1
}

func (g *Genre) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	//fmt.Println("UnmarshalJSON Genre: ", string(data))
	s := string(data)
	index1 := IndexOfSubstring(s, "\"")
	s = s[index1+1:]
	index2 := IndexOfSubstring(s, "\"")
	s = s[index2+1:]
	index3 := IndexOfSubstring(s, "\"")
	s = s[index3+1:]
	index4 := IndexOfSubstring(s, "\"")
	s = s[:index4]
	*g = Genre(s)
	// err := json.Unmarshal(data, &g)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// Movie описывает фильм
type Movie struct {
	Title  string  `json:"name"`
	Year   int     `json:"released_at"`
	Genres []Genre `json:"tags"`
}

// конец решения

func main() {
	const src = `{
		"name": "Interstellar",
		"released_at": 2014,
		"director": "Christopher Nolan",
		"tags": [
			{ "name": "Adventure" },
			{ "name": "Drama" },
			{ "name": "Science Fiction" }
		],
		"duration": "2h49m",
		"rating": "★★★★★"
	}`

	var m Movie
	err := json.Unmarshal([]byte(src), &m)
	fmt.Println(err)
	// nil
	fmt.Println(m)
	// {Interstellar 2014 [Adventure Drama Science Fiction]}
}
