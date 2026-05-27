package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no timezones with addresses were provided")
	}

	var wg sync.WaitGroup
	for _, arg := range os.Args[1:] {
		argSlice := strings.SplitN(arg, "=", -1)
		if len(argSlice) != 2 {
			log.Fatal("want an argument like zone=localhost:8000")
		}

		zone, addr := argSlice[0], argSlice[1]

		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		wg.Add(1)
		go func() {
			defer wg.Done()
			clockReader(zone, conn)
		}()
	}
	wg.Wait()
}

func clockReader(zone string, conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("%s: %s\n", zone, scanner.Text())
	}
}
