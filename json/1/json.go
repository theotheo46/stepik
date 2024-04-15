package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// начало решения

// Duration описывает продолжительность фильма
type Duration time.Duration

func process_duration_string(s string) string {

	if strings.HasPrefix(s, "0h") {
		s = strings.ReplaceAll(s, "0h", "")
	}
	if strings.HasSuffix(s, "0s") {
		s = strings.ReplaceAll(s, "0s", "")
	}
	if strings.HasSuffix(s, "h0m") {
		s = strings.ReplaceAll(s, "0m", "")
	}
	return s
}

func (d Duration) MarshalJSON() ([]byte, error) {
	s := time.Duration(d).String()
	b := make([]byte, 0, 10)
	b = append(b, '"')
	b = append(b, []byte(process_duration_string(s))...)
	b = append(b, '"')
	return b, nil
}

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, 10)

	b = append(b, '"')
	for i := 0; i < int(r); i++ {
		b = append(b, []byte("\u2605")...)
	}
	for i := 0; i < 5-int(r); i++ {
		b = append(b, []byte("\u2606")...)
	}
	b = append(b, '"')
	return b, nil
}

// Movie описывает фильм
type Movie struct {
	Title    string
	Year     int
	Director string
	Genres   []string
	Duration Duration
	Rating   Rating
}

// MarshalMovies кодирует фильмы в JSON.
//   - если indent = 0 - использует json.Marshal
//   - если indent > 0 - использует json.MarshalIndent
//     с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {

	if indent == 0 {
		b, err := json.Marshal(movies)
		if err != nil {
			return "", err
		}
		return string(b), nil
	} else {
		b, err := json.MarshalIndent(movies, "", strings.Repeat(" ", indent))
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

}

// конец решения

func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(2*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 36*time.Minute),
		Rating:   4,
	}

	s, err := MarshalMovies(4, m1, m2)
	fmt.Println(err)
	// nil
	fmt.Println(s)
	/*
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h36m",
		        "Rating": "★★★★☆"
		    }
		]
	*/
}
