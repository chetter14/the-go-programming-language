package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	var totalN sync.WaitGroup
	for _, root := range roots {
		totalN.Add(1)

		fileSizes := make(chan int64)
		var n sync.WaitGroup

		n.Add(1)
		go walkDir(root, &n, fileSizes)
		go printResults(root, &totalN, fileSizes)

		go func() {
			n.Wait()
			close(fileSizes)
		}()
	}

	totalN.Wait()
	fmt.Println("----- Completed! -----")
}

func printResults(root string, totalN *sync.WaitGroup, fileSizes <-chan int64) {
	defer totalN.Done()

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(root, nfiles, nbytes)
		}
	}
	printDiskUsage(root, nfiles, nbytes) // final totals
}

func printDiskUsage(root string, nfiles, nbytes int64) {
	fmt.Printf("%s: %d files %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
