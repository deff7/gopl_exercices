package main

import "testing"

func TestReverse(t *testing.T) {
	in := [size]int{}
	for i, _ := range in {
		in[i] = i
	}

	reverse(&in)

	for i, el := range in {
		exp := len(in) - i - 1
		if el != exp {
			t.Fatalf("element with index %d is %d but expected %d", i, el, exp)
		}
	}
}
