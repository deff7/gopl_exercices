package main

import (
	"bytes"
	"strings"
)

func commasFloat(in string) string {
	if len(in) <= 3 {
		return in
	}

	var (
		buf   bytes.Buffer
		comma = ""
		i     = 0
	)

	dotIdx := strings.IndexRune(in, '.')
	if dotIdx == -1 {
		dotIdx = len(in)
	}

	step := dotIdx % 3
	if step == 0 {
		step = 3
	}

	for i < dotIdx {
		buf.WriteString(comma)
		buf.WriteString(in[i : i+step])
		i += step
		step = 3
		comma = ","
	}

	if dotIdx != len(in) {
		buf.WriteRune('.')
		i++
	}

	comma = ""
	for i < len(in) {
		end := i + 3
		if end >= len(in) {
			end = len(in)
		}
		buf.WriteString(comma)
		buf.WriteString(in[i:end])
		i += 3
		comma = ","
	}

	return buf.String()
}
