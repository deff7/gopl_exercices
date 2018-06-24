package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

const httpPrefix = "http://"

func fetch(w io.Writer, url string) (code int) {
	if !strings.HasPrefix(url, httpPrefix) {
		url = httpPrefix + url
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
	return resp.StatusCode
}
