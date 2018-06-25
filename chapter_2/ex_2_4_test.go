// Run this test together with ex_2_3_test.go to avoid duplication
package main

import (
	"testing"

	"github.com/deff7/gopl_exercises/chapter_2/popcount"
)

func TestPopCountShift(t *testing.T) {
	commonTest(t, popcount.PopCountShift)
}

func BenchmarkPopCountShift(b *testing.B) {
	commonBench(b, popcount.PopCountShift)
}
