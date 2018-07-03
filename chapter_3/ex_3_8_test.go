package main

import (
	"image/color"
	"io/ioutil"
	"testing"

	"github.com/deff7/gopl_exercises/chapter_3/mandelbrot"
)

func bench(b *testing.B, f func(*mandelbrot.Options, complex128) color.Color) {
	for i := 0; i < b.N; i++ {
		opts.F = f
		mandelbrot.SetOptions(opts)
		mandelbrot.PNG(ioutil.Discard)
	}
}

func BenchmarkC128(b *testing.B) {
	bench(b, mandelbrot.MandelbrotC128)
}

func BenchmarkC64(b *testing.B) {
	bench(b, mandelbrot.MandelbrotC64)
}

func BenchmarkBigFloat(b *testing.B) {
	bench(b, mandelbrot.MandelbrotBigFloat)
}

func BenchmarkBigRat(b *testing.B) {
	bench(b, mandelbrot.MandelbrotBigRat)
}

var opts mandelbrot.Options

func init() {
	opts = mandelbrot.GetOptions()
	opts.Width = 32
	opts.Height = 32
	opts.Iterations = 10
}
