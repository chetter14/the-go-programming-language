package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	name    string
	outChan chan<- string // a client channel that receives messages
}

func (cl client) String() string {
	return cl.name
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages (from all the clients)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' message channels.
			for cli := range clients {
				cli.outChan <- msg
			}
		case cli := <-entering:
			var existingClients string
			if len(clients) > 0 {
				for curClient := range clients {
					existingClients += (curClient.name + " ")
				}
			} else {
				existingClients = "No clients connected yet"
			}
			cli.outChan <- existingClients
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.outChan)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 10) // a client channel that receives messages (add buffering for non-blocking operations)
	go clientWriter(conn, ch)

	// Ask a client to enter his name
	input := bufio.NewScanner(conn)

	ch <- "Enter your name: "
	if !input.Scan() {
		conn.Close()
		return
	}

	who := input.Text()
	if who == "" {
		who = conn.RemoteAddr().String()
	}

	ch <- "You are " + who
	messages <- who + " has arrived"

	curClient := client{name: who, outChan: ch}
	entering <- curClient

	timer := time.NewTimer(5 * time.Minute)
	stop := make(chan struct{})
	go func() {
		select {
		case <-timer.C:
			conn.Close()
		case <-stop:
			// do nothing (handleConn will call "conn.Close()")
		}
	}()

	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(5 * time.Minute)
	}
	// NOTE: ignoring potential errors from input.Err()

	close(stop)
	leaving <- curClient
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
