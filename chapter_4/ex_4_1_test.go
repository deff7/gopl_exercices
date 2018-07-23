package main

import (
	"crypto/sha256"
	"testing"
)

func TestBitDiff(t *testing.T) {
	for _, tc := range []struct {
		a, b, c byte
	}{
		{1, 1, 0},
		{0, 0, 0},
		{1, 2, 2},
		{1, 3, 1},
		{3, 1, 1},
		{2, 1, 2},
		{255, 254, 1},
		{255, 0, 8},
		{0, 255, 8},
		{253, 255, 1},
		{1, 255, 7},
	} {
		if r := bitDiff(tc.a, tc.b); r != tc.c {
			t.Errorf("expected diff between %b and %b is %d, but actual is %d", tc.a, tc.b, tc.c, r)
		}
	}
}

func TestDiffSHA256(t *testing.T) {
	var (
		a = "foo"
		b = "bar"
	)

	var (
		sa = sha256.Sum256([]byte(a))
		sb = sha256.Sum256([]byte(b))
	)

	var (
		bsa  = bitsString(sa[:])
		bsb  = bitsString(sb[:])
		diff = 0
	)

	for i := 0; i < len(bsa); i++ {
		if bsa[i] != bsb[i] {
			diff++
		}
	}

	if r := diffSHA256([]byte(a), []byte(b)); r != diff {
		t.Errorf("expected %d got %d", diff, r)
	}
}
