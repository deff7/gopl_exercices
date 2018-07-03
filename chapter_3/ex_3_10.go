package main

import "bytes"

func commas(in string) string {
	if len(in) <= 3 {
		return in
	}
	var (
		buf   bytes.Buffer
		step  = len(in) % 3
		i     = 0
		comma = ""
	)

	for i < len(in) {
		buf.WriteString(comma)
		buf.WriteString(in[i : i+step])
		i += step
		step = 3
		comma = ","
	}

	return buf.String()
}
