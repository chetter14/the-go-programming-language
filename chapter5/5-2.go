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
	elemsMap := make(map[string]int)
	visit(doc, elemsMap)

	fmt.Println(elemsMap)
}

func visit(n *html.Node, elemsMap map[string]int) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		elemsMap[n.Data]++
	}

	visit(n.FirstChild, elemsMap)
	visit(n.NextSibling, elemsMap)
}
