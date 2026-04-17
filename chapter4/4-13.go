package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

const OmdApiURL = "http://www.omdbapi.com/?i=tt3896198&apikey=86cf5b8d"

type Movie struct {
	Poster string `json:"Poster"` /* it's a poster url */
}

func main() {
	movieName := os.Args[1]

	resp, err := http.Get(OmdApiURL + "&t=" + movieName)
	if err != nil {
		log.Fatalf("Failed to GET url - %s\n", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("GET query failed: %s\n", resp.Status)
		resp.Body.Close()
		return
	}

	var movie Movie
	if err = json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		log.Fatalf("Failed to decode: %s\n", err)
		resp.Body.Close()
		return
	}

	resp.Body.Close()

	resp, err = http.Get(movie.Poster)
	if err != nil {
		log.Fatalf("Failed to GET movie poster url - %s\n", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("GET movie poster query failed: %s\n", resp.Status)
		resp.Body.Close()
		return
	}

	out, err := os.Create("poster.jpg")
	if err != nil {
		log.Fatalf("Failed to create new file: %s\n", err)
		resp.Body.Close()
		return
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("Failed to write image to file: %s\n", err)
		resp.Body.Close()
		return
	}

	resp.Body.Close()
}
