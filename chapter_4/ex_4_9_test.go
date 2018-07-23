package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestWordfreq(t *testing.T) {
	for _, tc := range []struct {
		in  string
		out map[string]int
	}{
		{"", map[string]int{}},
		{"abc a b abc", map[string]int{"abc": 2, "a": 1, "b": 1}},
	} {
		in := strings.NewReader(tc.in)
		if r := wordfreq(in); !reflect.DeepEqual(r, tc.out) {
			t.Errorf("wordfreq(%q)=%v, expected %v", tc.in, r, tc.out)
		}
	}
}
