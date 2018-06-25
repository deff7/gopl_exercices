package tempconv

import (
	"math"
	"testing"
)

func TestConversions(t *testing.T) {
	tc := Celsius(0)
	tk := Kelvin(0)
	tf := Fahrenheit(0)

	if CToK(tc) != FreezingK {
		t.Fail()
	}

	if KToC(tk) != AbsoluteZeroC {
		t.Fail()
	}

	if math.Abs(float64(FToK(tf))-255.37) > 1.0 {
		t.Fail()
	}
}
