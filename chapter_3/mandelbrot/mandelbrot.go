package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

type Options struct {
	Colorize   bool
	SSAA       bool
	Width      int
	Height     int
	Contrast   uint8
	Iterations uint8
	X, Y, Dev  float64
	F          func(*Options, complex128) color.Color
}

var options Options

func SetOptions(opts Options) {
	options = opts
}

func GetOptions() Options {
	return options
}

func mean(args ...uint8) uint8 {
	var r uint8
	for _, a := range args {
		r += a / uint8(len(args))
	}
	return r
}

func rgbaColors(args ...color.Color) []color.RGBA {
	cc := make([]color.RGBA, len(args))
	for i, c := range args {
		cc[i] = c.(color.RGBA)
	}
	return cc
}

func grayColors(args ...color.Color) []color.Gray {
	cc := make([]color.Gray, len(args))
	for i, c := range args {
		cc[i] = c.(color.Gray)
	}
	return cc
}

type imageSetter interface {
	image.Image
	Set(int, int, color.Color)
}

func PNG(w io.Writer) error {
	var (
		xmin, ymin, xmax, ymax = -2.0, -2.0, 2.0, 2.0
	)

	if options.Dev != 0 {
		var (
			x = options.X
			y = options.Y
			d = options.Dev
		)
		xmin = x - d
		xmax = x + d
		ymin = y - d
		ymax = y + d
	}

	var (
		width  = options.Width
		height = options.Height
	)

	if options.SSAA {
		width *= 2
		height *= 2
	}

	var img imageSetter

	if options.Colorize {
		img = image.NewRGBA(image.Rect(0, 0, width, height))
	} else {
		img = image.NewGray(image.Rect(0, 0, width, height))
	}
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			img.Set(px, py, options.F(&options, complex(x, y)))
		}
	}

	if options.SSAA {
		newImg := image.NewRGBA(image.Rect(0, 0, width/2, height/2))
		for py := 0; py < height; py += 2 {
			for px := 0; px < width; px += 2 {
				var c color.Color
				if options.Colorize {
					cc := rgbaColors(
						img.At(px, py),
						img.At(px, py+1),
						img.At(px+1, py),
						img.At(px+1, py+1),
					)
					var (
						r = mean(cc[0].R, cc[1].R, cc[2].R, cc[3].R)
						g = mean(cc[0].G, cc[1].G, cc[2].G, cc[3].G)
						b = mean(cc[0].B, cc[1].B, cc[2].B, cc[3].B)
					)
					c = color.RGBA{r, g, b, 0xFF}
				} else {
					cc := grayColors(
						img.At(px, py),
						img.At(px, py+1),
						img.At(px+1, py),
						img.At(px+1, py+1),
					)
					c = color.Gray{mean(cc[0].Y, cc[1].Y, cc[2].Y, cc[3].Y)}
				}
				newImg.Set(px/2, py/2, c)
			}
		}
		img = newImg
	}

	return png.Encode(w, img)
}

// X, Y, Width
func IntrestingPoint(i int) (x, y, w float64) {
	var pts = [][]float64{
		{-0.777807810193171, 0.131645108003206, 3.2e-6},
	}
	if i < 0 || i >= len(pts) {
		return
	}
	return pts[i][0], pts[i][1], pts[i][2]
}

var rootColors = []color.Color{
	color.RGBA{0xff, 0, 0, 0xff},
	color.RGBA{0, 0xff, 0, 0xff},
	color.RGBA{0, 0, 0xff, 0xff},
	color.RGBA{0xff, 0, 0xff, 0xff},
}

func init() {
	options = Options{
		Colorize:   false,
		Width:      1024,
		Height:     1024,
		Contrast:   contrast,
		Iterations: iterations,
		F:          MandelbrotC128,
	}
}
