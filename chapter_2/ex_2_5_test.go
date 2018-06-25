// Run this test together with ex_2_3_test.go to avoid duplication
package main

import (
	"testing"

	"github.com/deff7/gopl_exercises/chapter_2/popcount"
)

func TestPopCountRightmost(t *testing.T) {
	commonTest(t, popcount.PopCountRightmost)
}

func BenchmarkPopCountRightmost(b *testing.B) {
	commonBench(b, popcount.PopCountRightmost)
}
