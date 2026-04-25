package main

import (
	"golang.org/x/net/html"
)

func main() {
	ElementByID(nil, "val")
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil {
		if !post(n) {
			return false
		}
	}

	return true
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var found *html.Node
	pre := func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				/* Stop traveral */
				return false
			}
		}
		/* Continue traversal */
		return true
	}

	forEachNode(doc, pre, nil)
	return found
}
