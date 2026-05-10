package main

import (
	"flag"
	"fmt"
	"time"
)

var per = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *per)
	time.Sleep(*per)
	fmt.Println()
}
