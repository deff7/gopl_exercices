package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

var allowed = map[string]string{
	"a":      "href",
	"link":   "src",
	"img":    "src",
	"script": "src",
}

func getAttr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		for elem, link := range allowed {
			if n.Data == elem {
				links = append(links, getAttr(n, link))
			}
		}
	}

	links = visit(links, n.NextSibling)
	links = visit(links, n.FirstChild)

	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(visit([]string{}, doc))
}
