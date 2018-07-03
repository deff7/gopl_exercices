package main

import (
	"math"
	"os"

	"github.com/deff7/gopl_exercises/chapter_3/surface"
)

func main() {
	opts := surface.GetOptions()
	opts.Colorize = true
	surface.SetOptions(opts)
	surface.SVG(os.Stdout, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return math.Sin(r) * math.Cos(y/5)
	})
}
