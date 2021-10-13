package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	intervals "github.com/anicolaspp/intervals/interval"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		ints := genIntervals(5)
		merged := intervals.Merge(ints)
		io.WriteString(w, fmt.Sprintf("%s", ints))
		io.WriteString(w, fmt.Sprintf("%s", merged))
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func genIntervals(n int) []*intervals.Interval {
	result := []*intervals.Interval{}

	for i := 0; i <= n; i++ {
		x := rand.Intn(20)
		y := x + 10

		result = append(result, &intervals.Interval{x, y})
	}

	return result
}
