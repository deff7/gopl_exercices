package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/deff7/gopl_exercises/chapter_3/mandelbrot"
)

func main() {
	var (
		opts = mandelbrot.GetOptions()
		err  error
		ff   = make([]*os.File, 4)
	)

	for i := range ff {
		ff[i], err = os.Create(fmt.Sprintf("out_%d.png", i))
		checkErr(err)
		defer ff[i].Close()
	}

	//	opts.X, opts.Y, opts.Dev = mandelbrot.IntrestingPoint(0)
	opts.Colorize = false
	opts.Height = 512
	opts.Width = 512
	opts.Iterations = 100

	for i, f := range [](func(*mandelbrot.Options, complex128) color.Color){
		mandelbrot.MandelbrotC128,
		mandelbrot.MandelbrotC64,
		mandelbrot.MandelbrotBigFloat,
		mandelbrot.MandelbrotBigRat,
	} {
		opts.F = f
		mandelbrot.SetOptions(opts)
		err = mandelbrot.PNG(ff[i])
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
