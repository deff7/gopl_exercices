package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		magn int
		out  []int
	}{
		{[]int{}, 0, []int{}},
		{[]int{}, 1, []int{}},
		{[]int{}, -1, []int{}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3}, 1, []int{3, 1, 2}},
		{[]int{1, 2, 3}, -1, []int{2, 3, 1}},
		{[]int{1, 2, 3}, 2, []int{2, 3, 1}},
		{[]int{1, 2, 3}, -2, []int{3, 1, 2}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, -3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 5, []int{2, 3, 1}},
	} {
		r := make([]int, len(tc.in))
		copy(r, tc.in)
		rotate(r, tc.magn)
		if !reflect.DeepEqual(tc.out, r) {
			t.Errorf(
				"for slice %v and magnitude %d expected %v but got %v",
				tc.in, tc.magn, tc.out, r,
			)
		}
	}
}
