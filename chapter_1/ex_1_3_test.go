package main

import (
	"io/ioutil"
	"math/rand"
	"testing"
)

const inputSize = 50

var input []string

func BenchmarkEchoConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoConcat(ioutil.Discard, input)
	}
}

func BenchmarkEchoJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoJoin(ioutil.Discard, input)
	}
}

func init() {
	input = make([]string, inputSize)
	for i := range input {
		buf := make([]rune, inputSize)
		for j := range buf {
			buf[j] = 'a' + rand.Int31n('z'-'a')
		}
		input[i] = string(buf)
	}
}
