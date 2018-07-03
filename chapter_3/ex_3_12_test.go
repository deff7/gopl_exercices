package main

import "testing"

func TestAnagram(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		res  bool
	}{
		{"", "", false},
		{"123", "123", true},
		{"1234", "123", false},
		{"123", "321", true},
		{"123", "122", false},
		{"abc", "cab", true},
		{"абц", "бац", true},
	} {
		if r := anagram(tc.a, tc.b); r != tc.res {
			t.Errorf("expected %q and %q is anagram: %t", tc.a, tc.b, tc.res)
		}
	}
}
