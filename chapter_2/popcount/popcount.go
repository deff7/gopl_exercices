package popcount

var pc [256]byte

func init() {
	for i := 0; i < 256; i++ {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	res := 0
	var i uint
	for ; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

func PopCountShift(x uint64) int {
	res := 0
	for x != 0 {
		if x&1 == 1 {
			res++
		}
		x >>= 1
	}
	return res
}

func PopCountRightmost(x uint64) int {
	res := 0
	var prev uint64
	for x != 0 {
		prev = x
		x &= (x - 1)
		if prev != x {
			res++
		}
	}
	return res
}
