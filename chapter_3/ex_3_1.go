package main

import (
	"math"
	"os"

	"github.com/deff7/gopl_exercises/chapter_3/surface"
)

func main() {
	surface.SVG(os.Stdout, func(x, y float64) float64 {
		if x > 1 {
			return math.Inf(1)
		}
		return math.Sin(math.Hypot(x, y))
	})
}
