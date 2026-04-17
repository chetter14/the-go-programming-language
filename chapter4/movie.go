package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"release"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Action movie", Year: 2001, Color: false, Actors: []string{"Bob", "Alice"}},
	{Title: "Scary movie", Year: 2008, Color: true, Actors: []string{"John", "William", "Kate"}},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshalling failed - %s\n", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshalling failed - %s\n", err)
	}
	fmt.Println(titles)
}
