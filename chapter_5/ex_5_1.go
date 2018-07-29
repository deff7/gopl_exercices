package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func findlinks(r io.Reader) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, n.NextSibling)
	links = visit(links, n.FirstChild)
	return links
}

func main() {
	fmt.Println("vim-go")
}
