package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchBuffer(w io.Writer, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}

func fetch(w io.Writer, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
