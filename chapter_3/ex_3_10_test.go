package main

import "testing"

type testCase struct {
	in  string
	out string
}

var testCases []testCase

func TestComma(t *testing.T) {
	for _, tc := range testCases {
		if r := commas(tc.in); r != tc.out {
			t.Errorf("with %q expected %q but actual %q", tc.in, tc.out, r)
		}
	}
}

func init() {
	testCases = []testCase{
		{"", ""},
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"1234567", "1,234,567"},
	}
}
