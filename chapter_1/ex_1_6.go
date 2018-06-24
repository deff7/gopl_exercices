package main

import (
	"image/color"
	"os"

	"github.com/deff7/gopl_exercises/chapter_1/lissajous"
)

func main() {
	lissajous.SetPalette([]color.Color{
		color.Black,
		color.RGBA{0, 0xFF, 0, 0xFF},
		color.RGBA{0, 0xFF, 0xFF, 0xFF},
		color.RGBA{0xFF, 0, 0, 0xFF},
	})
	lissajous.GIF(
		os.Stdout,
		64,
		100,
		5,
		6,
		0.001,
	)
}
