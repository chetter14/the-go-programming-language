package main

import (
	"fmt"
	"io"
	"os"
)

type countingWriter struct {
	w   io.Writer
	cnt int64
}

func (c *countingWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.cnt += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriter{w: w}
	return cw, &cw.cnt
}

func main() {
	cw, n := CountingWriter(os.Stdout)
	fmt.Fprintln(cw, "hello, world")
	fmt.Println("bytes written:", *n)
}
