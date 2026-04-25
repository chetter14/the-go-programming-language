package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		visit(n.NextSibling)
		return
	}

	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	visit(n.FirstChild)
	visit(n.NextSibling)
}
