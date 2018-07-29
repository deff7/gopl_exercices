package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func visit(count map[string]int, n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	visit(count, n.NextSibling)
	visit(count, n.FirstChild)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	count := map[string]int{}
	visit(count, doc)
	fmt.Println(count)
}
