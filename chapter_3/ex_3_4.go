package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/deff7/gopl_exercises/chapter_3/surface"
)

var defaultOptions = surface.GetOptions()

func parseBool(dest *bool, param string, req *http.Request) (err error) {
	if v := req.FormValue(param); v != "" {
		*dest, err = strconv.ParseBool(v)
	}
	return
}

func parseInt(dest *int, param string, req *http.Request) (err error) {
	if v := req.FormValue(param); v != "" {
		*dest, err = strconv.Atoi(v)
	}
	return
}

func parseFloat(dest *float64, param string, req *http.Request) (err error) {
	if v := req.FormValue(param); v != "" {
		*dest, err = strconv.ParseFloat(v, 64)
	}
	return
}

func appendErr(err error) (errs []error) {
	if err != nil {
		errs = append(errs, err)
	}
	return
}

func handler(w http.ResponseWriter, req *http.Request) {
	var (
		opts = defaultOptions
		errs = []error{}
	)
	errs = appendErr(parseBool(&opts.SkipNaN, "skip_nan", req))
	errs = appendErr(parseBool(&opts.Colorize, "colorize", req))

	errs = appendErr(parseInt(&opts.Width, "width", req))
	errs = appendErr(parseInt(&opts.Height, "height", req))
	errs = appendErr(parseInt(&opts.Cells, "cells", req))

	errs = appendErr(parseFloat(&opts.XYRange, "range", req))

	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		for _, err := range errs {
			fmt.Fprintln(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	surface.SetOptions(opts)
	surface.SVG(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return math.Sin(r) * math.Cos(x+y)
	})
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
