package main

import (
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	reader io.Reader
	limit  int64
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	if lr.limit <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > lr.limit {
		p = p[:lr.limit]
	}
	n, err = lr.reader.Read(p)
	lr.limit -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}

func main() {
	r := LimitReader(strings.NewReader("hello world, this is a long string"), 5)
	b, _ := io.ReadAll(r)
	fmt.Printf("%q\n", b) // "hello"
}
