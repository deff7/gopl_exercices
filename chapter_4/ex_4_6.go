package main

import "bytes"

func stripSpaces(in string) string {
	var (
		prev rune
		buf  bytes.Buffer
	)

	for _, c := range in {
		if prev != 0 && prev == c {
			continue
		}
		prev = c
		buf.WriteRune(c)
	}
	return buf.String()
}
