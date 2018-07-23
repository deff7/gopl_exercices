package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
)

func bitDiff(a, b byte) (res byte) {
	for a > 0 || b > 0 {
		res += (a & 1) ^ (b & 1)
		a >>= 1
		b >>= 1
	}
	return
}

func bitsString(a []byte) string {
	var buf bytes.Buffer
	for _, b := range a {
		buf.WriteString(fmt.Sprintf("%08b", b))
	}
	return buf.String()
}

func diffSHA256(a, b []byte) (res int) {
	var (
		sa = sha256.Sum256(a)
		sb = sha256.Sum256(b)
	)

	for i := 0; i < len(sa); i++ {
		res += int(bitDiff(sa[i], sb[i]))
	}

	return
}

func main() {
	diff := diffSHA256([]byte(os.Args[1]), []byte(os.Args[2]))
	fmt.Println(strconv.Itoa(diff))
}
