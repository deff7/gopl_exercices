package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const httpPrefix = "http://"

func fetch(ch chan<- string, url string) {
	start := time.Now()
	if !strings.HasPrefix(url, httpPrefix) {
		url = httpPrefix + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("http.Get error: %v", err)
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(url[len(httpPrefix):] + ".tmp")
	if err != nil {
		ch <- fmt.Sprintf("os.Create error: %v", err)
		return
	}
	defer f.Close()

	io.Copy(f, resp.Body)

	ch <- fmt.Sprintf(
		"get %s with status code %d time elapsed: %.2fs",
		url,
		resp.StatusCode,
		time.Since(start).Seconds(),
	)
}

func main() {
	ch := make(chan string)

	start := time.Now()
	for _, url := range os.Args[1:] {
		go fetch(ch, url)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("total time elapsed: %.2fs\n", time.Since(start).Seconds())
}
