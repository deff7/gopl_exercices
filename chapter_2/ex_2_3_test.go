package main

import (
	"testing"

	"github.com/deff7/gopl_exercises/chapter_2/popcount"
)

type testCase struct {
	in  uint64
	out int
}

var testCases = []testCase{
	{0, 0},
	{1, 1},
	{^uint64(0) - 1, 63},
	{^uint64(0), 64},
}

func commonTest(t *testing.T, f func(uint64) int) {
	for _, tc := range testCases {
		if r := f(tc.in); r != tc.out {
			t.Errorf("with x=%d expected %d but have %d", tc.in, tc.out, r)
		}
	}
}

func commonBench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(^uint64(0))
	}
}

func TestPopCount(t *testing.T) {
	commonTest(t, popcount.PopCount)
}

func BenchmarkPopCount(b *testing.B) {
	commonBench(b, popcount.PopCount)
}

func TestPopCountLoop(t *testing.T) {
	commonTest(t, popcount.PopCountLoop)
}

func BenchmarkPopCountLoop(b *testing.B) {
	commonBench(b, popcount.PopCountLoop)
}
