package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, strings.Repeat("test abc", 100))
	}))
}

func commonBench(b *testing.B, f func(w io.Writer, url string)) {
	ts := testServer()
	defer ts.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(ioutil.Discard, ts.URL)
	}
	b.StopTimer()
}

func BenchmarkFetchBuffer(b *testing.B) {
	commonBench(b, fetchBuffer)
}

func BenchmarkFetch(b *testing.B) {
	commonBench(b, fetch)
}
