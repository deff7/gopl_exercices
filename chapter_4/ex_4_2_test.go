package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"reflect"
	"testing"
)

func TestSHA(t *testing.T) {
	var (
		in   = []byte("foo")
		buf  bytes.Buffer
		s256 = sha256.Sum256(in)
		s384 = sha512.Sum384(in)
		s512 = sha512.Sum512(in)
	)

	for _, tc := range []struct {
		mode    string
		out     []byte
		isError bool
	}{
		{"SHA256", s256[:], false},
		{"SHA384", s384[:], false},
		{"SHA512", s512[:], false},
		{"SHA666", []byte{}, false},
	} {
		buf.Reset()
		err := sha(&buf, in, tc.mode)
		if tc.isError && err == nil {
			t.Error("expect error")
		}
		if !tc.isError && !reflect.DeepEqual(tc.out, buf.Bytes()) {
			t.Errorf("expected %q to equal %q", buf.Bytes(), tc.out)
		}
	}

}
