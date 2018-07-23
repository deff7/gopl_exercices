package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func sha(w io.Writer, in []byte, mode string) (err error) {
	switch mode {
	case "SHA256":
		b := sha256.Sum256(in)
		w.Write(b[:])
	case "SHA384":
		b := sha512.Sum384(in)
		w.Write(b[:])
	case "SHA512":
		b := sha512.Sum512(in)
		w.Write(b[:])
	default:
		err = fmt.Errorf("unknown algorithm %q", mode)
	}
	return
}

func main() {
	mode := flag.String("mode", "SHA256", "set SHA-2 algorithm, available:\nSHA256\nSHA384\nSHA512")
	flag.Parse()
	if err := sha(os.Stdout, []byte(os.Args[1]), *mode); err != nil {
		log.Fatal(err)
	}
}
