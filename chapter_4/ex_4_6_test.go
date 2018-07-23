package main

import (
	"testing"
)

func TestStripSpaces(t *testing.T) {
	for _, tc := range []struct {
		in, out string
	}{
		{"", ""},
		{"abc", "abc"},
		{" abc", " abc"},
		{"abc ", "abc "},
		{"abc  ", "abc "},
		{"a\t\tbc", "a\tbc"},
		{"  \t\t\n\n\v\v\f\f\r\r\u0085\u0085\u00a0\u00a0", " \t\n\v\f\r\u0085\u00a0"},
		{"a   b", "a b"},
		{"юни  код", "юни код"},
	} {
		if r := stripSpaces(tc.in); tc.out != r {
			t.Errorf("with %q expect %q got %q", tc.in, tc.out, r)
		}
	}
}
