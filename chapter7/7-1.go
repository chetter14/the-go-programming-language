package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	wordScanner := bufio.NewScanner(bytes.NewReader(p))
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		(*w)++
	}
	return len(p), nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	lineScanner := bufio.NewScanner(bytes.NewReader(p))
	for lineScanner.Scan() {
		(*l)++
	}
	return len(p), nil
}

func main() {
	var wc WordCounter
	fmt.Fprintf(&wc, "hello world\n")
	fmt.Println(wc)

	wc = 0
	fmt.Fprintf(&wc, "it is a bunch of words of len 9")
	fmt.Println(wc)

	var lc LineCounter
	fmt.Fprintf(&lc, "hello world\n")
	fmt.Println(lc)

	lc = 0
	fmt.Fprintf(&lc, "it is \n a bunch of words \n of line \n count 4")
	fmt.Println(lc)
}
