package main

import (
	"log"
	"os"

	"github.com/deff7/gopl_exercises/chapter_3/mandelbrot"
)

func main() {
	opts := mandelbrot.GetOptions()
	opts.Colorize = true
	mandelbrot.SetOptions(opts)
	err := mandelbrot.PNG(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
