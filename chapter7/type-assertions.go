package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f, ok := w.(*os.File)
	fmt.Println(f, ok)

	c, ok := w.(*bytes.Buffer)
	fmt.Println(c, ok)

	rw, ok := w.(io.ReadWriter)
	fmt.Println(rw, ok)

	w = new(bytes.Buffer)
	rw, ok = w.(io.ReadWriter)
	fmt.Println(rw, ok)

	fmt.Println("End")
}
