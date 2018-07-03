package surface

import (
	"fmt"
	"image/color"
	"io"
	"math"
	"strings"
)

type Options struct {
	SkipNaN         bool
	Colorize        bool
	Width, Height   int
	Cells           int
	XYRange         float64
	XYScale, ZScale float64
	Angle           float64
}

var options Options
var indicies [][]int
var sin30, cos30 = math.Sin(math.Pi / 6), math.Cos(math.Pi / 6)

func GetOptions() Options {
	return options
}

func SetOptions(opts Options) {
	options = opts
	if opts.XYRange*opts.ZScale == 0 {
		calculateScales(&options)
	}
}

var (
	valleyColor = color.RGBA{0, 0, 0xFF, 0}
	peakColor   = color.RGBA{0xFF, 0, 0, 0}
)

func colorHeight(z, min, max float64) color.Color {
	var (
		r, g, b uint8
		t       = (z - min) / (max - min)
	)

	interpolate := func(ai, bi uint8, t float64) uint8 {
		var (
			a = float64(ai)
			b = float64(bi)
		)
		return uint8(a + (b-a)*t)
	}

	r = interpolate(valleyColor.R, peakColor.R, t)
	g = interpolate(valleyColor.G, peakColor.G, t)
	b = interpolate(valleyColor.B, peakColor.B, t)

	return color.RGBA{r, g, b, 1}
}

type polygon struct {
	Points []float64
	Z      float64
}

func SVG(out io.Writer, f func(x, y float64) float64) {
	var (
		polygons = []polygon{}
		min, max float64
	)
	for i := 0; i < options.Cells; i++ {
		for j := 0; j < options.Cells; j++ {
			var (
				err  error
				poly = polygon{make([]float64, 0, 8), 0.0}
				z    = 0.0
			)

			for _, idx := range indicies {
				var x, y float64
				x, y, err = corner(i+idx[0], j+idx[1], f, &min, &max, &z)
				if err != nil {
					break
				}
				poly.Points = append(poly.Points, x, y)
			}
			if err != nil {
				continue
			}
			poly.Z = z / 8
			polygons = append(polygons, poly)
		}
	}

	writeHeader(out)
	for _, poly := range polygons {
		writePolygon(out, poly, min, max)
	}
	writeFooter(out)
}

func corner(i, j int, f func(float64, float64) float64, min, max, fz *float64) (float64, float64, error) {
	x := options.XYRange * (float64(i)/float64(options.Cells) - 0.5)
	y := options.XYRange * (float64(j)/float64(options.Cells) - 0.5)
	z := f(x, y)

	if *min > z {
		*min = z
	}
	if *max < z {
		*max = z
	}
	*fz += z

	if options.SkipNaN && (math.IsNaN(z) || math.IsInf(z, 0)) {
		return 0, 0, fmt.Errorf("NaN")
	}

	sx := float64(options.Width/2) + (x-y)*cos30*options.XYScale
	sy := float64(options.Height/2) + (x+y)*sin30*options.XYScale - z*options.ZScale
	return sx, sy, nil
}

func writeHeader(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: black; fill: white; stroke-width: 0.7;' "+
		"width='%d' height='%d'>", options.Width, options.Height)
}

func writePolygon(out io.Writer, poly polygon, min, max float64) {
	p := poly.Points
	t := strings.TrimSpace(strings.Repeat("%g,%g ", len(p)/2))

	args := make([]interface{}, len(p))
	for i, v := range p {
		args[i] = v
	}

	templ := fmt.Sprintf("<polygon points='%s'", t)
	if options.Colorize {
		r, g, b, _ := colorHeight(poly.Z, min, max).RGBA()
		c := fmt.Sprintf("#%02x%02x%02x", uint8(r), uint8(g), uint8(b))
		templ += fmt.Sprintf(" style='fill:%s'", c)
	}
	templ += "/>\n"
	fmt.Fprintf(out, templ, args...)
}

func writeFooter(out io.Writer) {
	fmt.Fprintf(out, "</svg>")
}

func calculateScales(opts *Options) {
	opts.XYScale = float64(opts.Width) / 2.0 / opts.XYRange
	opts.ZScale = float64(opts.Height) * 0.4
}

func init() {
	indicies = [][]int{
		[]int{1, 0},
		[]int{0, 0},
		[]int{0, 1},
		[]int{1, 1},
	}
	options = Options{
		SkipNaN:  true,
		Colorize: false,
		Width:    600,
		Height:   320,
		Cells:    100,
		XYRange:  15.0,
		Angle:    math.Pi / 6,
	}
	calculateScales(&options)
}
