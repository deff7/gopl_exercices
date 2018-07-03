package main

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/deff7/gopl_exercises/chapter_3/mandelbrot"
)

func parseFloat(req *http.Request, param string, dest *float64) (err error) {
	if v := req.FormValue(param); v != "" {
		var f float64
		f, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return
		}
		*dest = f
	}
	return nil
}

func appendErrs(errs []error, err error) []error {
	if err == nil {
		return errs
	}
	return append(errs, err)
}

func handler(w http.ResponseWriter, req *http.Request) {
	var (
		errs = []error{}
		opts = mandelbrot.GetOptions()
	)

	errs = appendErrs(errs, parseFloat(req, "x", &opts.X))
	errs = appendErrs(errs, parseFloat(req, "y", &opts.Y))

	oldDev := opts.Dev
	errs = appendErrs(errs, parseFloat(req, "zoom", &opts.Dev))
	if oldDev != opts.Dev {
		if opts.Dev == 0 {
			errs = appendErrs(errs, errors.New("zoom can't be zero"))
		} else {
			opts.Dev = 1.0 / opts.Dev
		}
	}

	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		var buf bytes.Buffer
		for _, e := range errs {
			buf.WriteString(e.Error())
			buf.WriteRune('\n')
		}
		w.Write(buf.Bytes())
		return
	}
	mandelbrot.SetOptions(opts)
	mandelbrot.PNG(w)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
