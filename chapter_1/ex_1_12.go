package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/deff7/gopl_exercises/chapter_1/lissajous"
)

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var (
		nframes, _ = strconv.Atoi(q.Get("nframes"))
		size, _    = strconv.Atoi(q.Get("size"))
		cycles, _  = strconv.Atoi(q.Get("cycles"))
		delay, _   = strconv.Atoi(q.Get("delay"))
		res, _     = strconv.ParseFloat(q.Get("res"), 64)
	)
	lissajous.GIF(w, nframes, size, cycles, delay, res)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
