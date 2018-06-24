package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black}

func SetPalette(newPalette []color.Color) {
	palette = newPalette
}

func GIF(out io.Writer, nframes, size, cycles, delay int, res float64) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < math.Pi*float64(2*cycles); t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			var blackIndex uint8 = uint8(1 + (i % (len(palette) - 1)))
			img.SetColorIndex(
				size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5),
				blackIndex,
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
