package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(205)
	}))
	defer ts.Close()

	if fetch(ioutil.Discard, ts.URL) != 205 {
		t.Error("must return status code")
	}
}
