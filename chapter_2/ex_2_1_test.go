package main

import (
	"math"
	"testing"

	"github.com/deff7/gopl_exercises/chapter_2/tempconv"
)

func TestConversions(t *testing.T) {
	tc := tempconv.Celsius(0)
	tk := tempconv.Kelvin(0)
	tf := tempconv.Fahrenheit(0)

	if tempconv.CToK(tc) != tempconv.FreezingK {
		t.Fail()
	}

	if tempconv.KToC(tk) != tempconv.AbsoluteZeroC {
		t.Fail()
	}

	if math.Abs(float64(tempconv.FToK(tf))-255.37) > 1.0 {
		t.Fail()
	}
}
