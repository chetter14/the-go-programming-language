package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func main() {
	CountWordsAndImages("golang.org")
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	w, i := countWordsAndImages(n.FirstChild)
	words += w
	images += i

	w, i = countWordsAndImages(n.NextSibling)
	words += w
	images += i

	return
}
