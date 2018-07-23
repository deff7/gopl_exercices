package main

import "unicode/utf8"

func reverseUTF8(in []byte) {
	buf := make([]byte, len(in))
	i := 0
	for i < len(in) {
		_, n := utf8.DecodeRune(in[i:])
		copy(buf[len(in)-(i+n):len(in)-i], in[i:i+n])
		i += n
	}
	copy(in, buf)
}

func reverseMem(in []byte) {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}

	last := 0
	for i := 0; i < len(in); i++ {
		if utf8.RuneStart(in[i]) {
			for s, e := last, i; s < e; s, e = s+1, e-1 {
				in[s], in[e] = in[e], in[s]
			}
			last = i + 1
		}
	}
}
