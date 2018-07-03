package mandelbrot

import (
	"image/color"
	"math/big"
	"math/cmplx"
)

const (
	iterations = 200
	contrast   = 15
)

func MandelbrotC128(_ *Options, z complex128) color.Color {
	var (
		v complex128
	)
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			if options.Colorize {
				return color.RGBA{n, 255 - n, 50 % ((n + 1) * 3), 255}
			} else {
				return color.Gray{255 - contrast*n}
			}
		}
	}
	return color.Black
}

func MandelbrotC64(_ *Options, _z complex128) color.Color {
	var (
		v complex64
		z = complex64(_z)
	)
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			if options.Colorize {
				return color.RGBA{n, 255 - n, 50 % ((n + 1) * 3), 255}
			} else {
				return color.Gray{255 - contrast*n}
			}
		}
	}
	return color.Black
}

func MandelbrotBigFloat(_ *Options, z complex128) color.Color {
	var (
		real  = big.NewFloat(real(z))
		imag  = big.NewFloat(imag(z))
		vreal = big.NewFloat(0)
		vimag = big.NewFloat(0)
		two   = big.NewFloat(2)
		four  = big.NewFloat(4)
	)

	bf := func() *big.Float {
		return &big.Float{}
	}
	for n := uint8(0); n < iterations; n++ {
		vrSquare := bf().Mul(vreal, vreal)
		viSquare := bf().Mul(vimag, vimag)
		vrvi := bf().Mul(vreal, vimag)

		vreal.Add(bf().Sub(vrSquare, viSquare), real)
		vimag.Add(bf().Mul(two, vrvi), imag)
		if bf().Add(vrSquare, viSquare).Cmp(four) > 0 {
			if options.Colorize {
				return color.RGBA{n, 255 - n, 50 % ((n + 1) * 3), 255}
			} else {
				return color.Gray{255 - contrast*n}
			}
		}
	}
	return color.Black
}

func MandelbrotBigRat(_ *Options, z complex128) color.Color {
	br := func(f float64) *big.Rat {
		return (&big.Rat{}).SetFloat64(f)
	}

	var (
		real  = br(real(z))
		imag  = br(imag(z))
		vreal = br(0)
		vimag = br(0)
		two   = br(2)
		four  = br(4)
	)

	for n := uint8(0); n < iterations; n++ {
		vrSquare := br(0).Mul(vreal, vreal)
		viSquare := br(0).Mul(vimag, vimag)
		vrvi := br(0).Mul(vreal, vimag)

		vreal.Add(br(0).Sub(vrSquare, viSquare), real)
		vimag.Add(br(0).Mul(two, vrvi), imag)
		if br(0).Add(vrSquare, viSquare).Cmp(four) > 0 {
			if options.Colorize {
				return color.RGBA{n, 255 - n, 50 % ((n + 1) * 3), 255}
			} else {
				return color.Gray{255 - contrast*n}
			}
		}
	}
	return color.Black
}
