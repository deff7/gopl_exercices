package main

import (
	"bytes"
	"io"
	"math"
	"testing"

	"github.com/deff7/gopl_exercises/chapter_3/surface"
)

func generateSvg(w io.Writer) {
	surface.SVG(w, func(x, y float64) float64 {
		if x > 1 {
			return math.Inf(1)
		}
		return math.Sin(math.Hypot(x, y))
	})
}

func TestInvalidPolygons(t *testing.T) {
	var (
		buf1 = new(bytes.Buffer)
		buf2 = new(bytes.Buffer)
	)

	// svg without invalid polygons
	generateSvg(buf1)

	opts := surface.GetOptions()
	opts.SkipNaN = false
	surface.SetOptions(opts)
	// svg with invalid polygons
	generateSvg(buf2)

	if buf1.Len() >= buf2.Len() {
		t.Error("generator with option SkipNaN set to true must skip invalid polygons and decrease SVG size")
	}
}
