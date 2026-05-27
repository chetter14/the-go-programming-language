package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{}
		// signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	tcpConn := conn.(*net.TCPConn)
	tcpConn.CloseWrite()

	<-done
	// wait for background goroutine to finish
}
