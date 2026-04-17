package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const XkcdURL = "https://xkcd.com/"

type ComicTranscript struct {
	Transcript string
}

func main() {
	for _, val := range os.Args[1:] {
		url := XkcdURL + val + "/info.0.json"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to GET url - %s", err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("GET query failed: %s\n", resp.Status)
			resp.Body.Close()
			continue
		}

		var result ComicTranscript
		if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Fatalf("Failed to decode transcript - %s\n", err)
			resp.Body.Close()
			continue
		}

		log.Printf("url:\t%s\n", url)
		log.Printf("transcript:\t%s\n", result.Transcript)

		resp.Body.Close()
		fmt.Println()
	}
}
