package main

import (
	"image/color"
	"log"
	"math/cmplx"
	"os"

	"github.com/deff7/gopl_exercises/chapter_3/mandelbrot"
)

func newton(opts *mandelbrot.Options, z complex128) color.Color {
	roots := []complex128{
		1,
		-1,
		-1i,
		1i,
	}

	var (
		contrast   = opts.Contrast
		iterations = opts.Iterations
	)

	for n := uint8(0); n < iterations; n++ {
		z -= (z*z*z*z - 1) / (4 * z * z * z)
		for i, r := range roots {
			if cmplx.Abs(z-r) < 0.000001 {
				if opts.Colorize {
					switch i {
					case 0:
						return color.RGBA{n * contrast, 0, 0, 0xff}
					case 1:
						return color.RGBA{0, n * contrast, 0, 0xff}
					case 2:
						return color.RGBA{0, 0, n * contrast, 0xff}
					default:
						return color.RGBA{n * contrast, 0, n * contrast, 0xff}
					}
				} else {
					return color.Gray{n * contrast}
				}
			}
		}
	}
	return color.White
}

func main() {
	opts := mandelbrot.GetOptions()
	opts.Colorize = true
	opts.F = newton
	mandelbrot.SetOptions(opts)
	err := mandelbrot.PNG(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
