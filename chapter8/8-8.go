package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	receivedMsg := make(chan struct{}, 1)
	done := make(chan struct{})

	go func() {
		timer := time.NewTimer(10 * time.Second)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				c.Close()
				return
			case <-receivedMsg:
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(10 * time.Second)
			case <-done:
				return
			}
		}
	}()

	for input.Scan() {
		select {
		case receivedMsg <- struct{}{}:
		default:
		}
		echo(c, input.Text(), 1*time.Second)
	}
	close(done)
	c.Close()
}
