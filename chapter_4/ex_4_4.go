package main

func rotate(in []int, magn int) {
	if len(in) == 0 {
		return
	}

	magn = magn % len(in)
	if magn == 0 {
		return
	}
	if magn < 0 {
		magn = len(in) + magn
	}

	for i, _ := range in {
		j := (i + magn) % len(in)
		in[0], in[j] = in[j], in[0]
	}
}
