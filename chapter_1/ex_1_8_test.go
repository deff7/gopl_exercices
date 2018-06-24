package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "yahoy!")
	}))
	defer ts.Close()

	url := ts.URL[7:]

	buf := new(bytes.Buffer)
	fetch(buf, url)
	if buf.String() != "yahoy!" {
		t.Error("must set http prefix to url")
	}
}
