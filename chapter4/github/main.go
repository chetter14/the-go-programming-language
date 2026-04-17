package main

import (
	"fmt"
	"gh/github"
	"log"
	"os"
)

func main() {
	res, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", res.TotalCount)
	for _, item := range res.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
