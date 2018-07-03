package main

import "testing"

const (
	KB = 1 << (10 + 10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func TestConstants(t *testing.T) {
	assert := func(a, b float64) {
		if a != b {
			t.Errorf("expect %f is equal %f", a, b)
		}
	}
	assert(KB, 1024)
	assert(MB, 1048576)
	assert(GB, 1073741824)
	assert(TB, 1099511627776)
	assert(PB, 1125899906842624)
	assert(EB, 1152921504606846976)
	assert(ZB, 1180591620717411303424)
	assert(YB, 1208925819614629174706176)
}
