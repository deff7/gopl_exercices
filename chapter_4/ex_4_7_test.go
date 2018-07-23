package main

import (
	"reflect"
	"testing"
)

func commonTest(t *testing.T, f func([]byte)) {
	for _, tc := range []struct {
		in, out []byte
	}{
		{[]byte("abc"), []byte("cba")},
		{[]byte("абц"), []byte("цба")},
		{[]byte("abcйцук"), []byte("куцйcba")},
	} {
		in := make([]byte, len(tc.in))
		copy(in, tc.in)
		f(in)
		if !reflect.DeepEqual(in, tc.out) {
			t.Errorf("reverseSimple(%q) = %q, expected %q", tc.in, in, tc.out)
		}
	}
}

func TestReverseUTF8(t *testing.T) {
	commonTest(t, reverseUTF8)
}

func TestReverseMem(t *testing.T) {
	commonTest(t, reverseMem)
}

var input = []byte("abc шёл")

func commonBench(b *testing.B, f func([]byte)) {
	for i := 0; i < b.N; i++ {
		f(input)
	}
}

func BenchmarkReverseUTF8(b *testing.B) {
	commonBench(b, reverseUTF8)
}

func BenchmarkReverseMem(b *testing.B) {
	commonBench(b, reverseMem)
}
