package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	result := make(chan string)
	for _, arg := range os.Args[1:] {
		go fetch(arg, result)
	}

	select {
	case str := <-result:
		log.Printf("received a result - %s", str)
	case <-done:
		log.Printf("received a done")
	}

	close(done)
	log.Printf("----- Completed! -----")
}

func fetch(url string, result chan<- string) {
	if cancelled() {
		log.Printf("do not fetch url - %s", url)
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("failed to create a request: %s", url)
		return
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("failed to fetch url %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body")
		return
	}

	select {
	case result <- url:
	case <-done:
	}
}
