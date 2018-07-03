package main

import "testing"

func TestCommas(t *testing.T) {
	for _, tc := range []struct {
		in, out string
	}{
		{"", ""},
		{"1", "1"},
		{"123", "123"},
		{"12345678", "12,345,678"},
		{"1.0", "1.0"},
		{"1.", "1."},
		{"-1.0", "-1.0"},
		{"123.456", "123.456"},
		{"123.5678", "123.567,8"},
		{"1234.5678", "1,234.567,8"},
		{"-1234.5678", "-1,234.567,8"},
	} {
		if r := commasFloat(tc.in); r != tc.out {
			t.Errorf("with %q expected %q, got %q", tc.in, tc.out, r)
		}
	}
}
