package main

import (
	"reflect"
	"testing"
)

func TestCharcount(t *testing.T) {
	for _, tc := range []struct {
		in  string
		out countMap
	}{
		{"", countMap{}},
		{"aa1\u0362\b", countMap{
			"letter": {'a': 2},
			"digit":  {'1': 1},
			"mark":   {'\u0362': 1},
			"other":  {'\b': 1},
		}},
	} {
		if r := charcount(tc.in); !reflect.DeepEqual(r, tc.out) {
			t.Errorf("charcount(%q) = %v, expected %v", tc.in, r, tc.out)
		}
	}
}
