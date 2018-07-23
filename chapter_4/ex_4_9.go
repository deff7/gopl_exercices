package main

import (
	"bufio"
	"io"
)

func wordfreq(r io.Reader) map[string]int {
	var (
		s   = bufio.NewScanner(r)
		res = map[string]int{}
	)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		res[s.Text()]++
	}

	return res
}
