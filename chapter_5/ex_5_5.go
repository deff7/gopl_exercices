package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func visit(n *html.Node, f func(n *html.Node)) {
	if n == nil {
		return
	}
	f(n)
	visit(n.NextSibling, f)
	visit(n.FirstChild, f)
}

func countWordsAndImages(doc *html.Node) (words, images int) {
	visit(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			images++
		}

		if n.Type == html.TextNode {
			s := bufio.NewScanner(strings.NewReader(n.Data))
			s.Split(bufio.ScanWords)
			for s.Scan() {
				words++
			}
			if err := s.Err(); err != nil {
				log.Println(err)
			}
		}
	})
	return
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func main() {
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("words: %d\nimages: %d\n", words, images)
}
