package main

import "testing"

func TestSnakeCase(t *testing.T) {
	for _, tc := range []struct {
		in, out string
	}{
		{"", ""},
		{" ", "_"},
		{"a b c", "a_b_c"},
		{"A B C", "a_b_c"},
		{"Abc FoO BAR", "abc_foo_bar"},
	} {
		if r := snakeCase(tc.in); r != tc.out {
			t.Errorf("snakeCase(%q)=%q want %q", tc.in, r, tc.out)
		}
	}
}
