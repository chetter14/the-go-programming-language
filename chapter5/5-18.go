package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	if filename, size, err := fetch("https://golang.org"); err != nil {
		fmt.Printf("failed to fetch: %v\n", err)
	} else {
		fmt.Printf("filename - %s, size - %d bytes\n", filename, size)
	}
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	fmt.Println(local)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
