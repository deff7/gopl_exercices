package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestCountLines(t *testing.T) {
	var (
		in      = []string{"123\nabc\nabc", "foo\nfoo\n123"}
		counts  = make(map[string]int)
		foundIn = make(map[string]map[string]bool)
	)

	for i, s := range in {
		r := strings.NewReader(s)
		countLines(r, strconv.Itoa(i), counts, foundIn)
	}

	if len(counts) != 3 {
		t.Error("must count all lines")
	}

	if counts["123"] != 2 {
		t.Error("must couts lines in different readers")
	}

	if counts["abc"] != 2 {
		t.Error("must count lines across one reader")
	}
}
