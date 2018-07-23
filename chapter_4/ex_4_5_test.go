package main

import (
	"reflect"
	"testing"
)

func TestUniqAdjacentStrings(t *testing.T) {
	for _, tc := range []struct {
		in, out []string
	}{
		{[]string{}, []string{}},
		{[]string{""}, []string{""}},
		{[]string{"a", "a"}, []string{"a"}},
		{[]string{"b", "a", "a"}, []string{"b", "a"}},
		{[]string{"a", "b", "a"}, []string{"a", "b", "a"}},
	} {
		if r := uniqAdjacentStrings(tc.in); !reflect.DeepEqual(tc.out, r) {
			t.Errorf("expected that uniq strings of %q is %q, but got %q", tc.in, tc.out, r)
		}
	}
}
