package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func visit(n *html.Node) {
	if n == nil {
		return
	}

	switch n.Type {
	case html.TextNode:
		fmt.Println(n.Data)
	case html.ElementNode:
		if n.Data == "style" || n.Data == "script" {
			return
		}
	}
	visit(n.NextSibling)
	visit(n.FirstChild)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	visit(doc)
}
