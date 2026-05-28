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
	}()

	mustCopy(conn, os.Stdin)

	if tcpConn, ok := conn.(*net.TCPConn); ok {
		log.Println("Closed the tcp connection")
		tcpConn.CloseWrite() // signal server: no more input
	}

	<-done // wait for server to finish echoing, then EOF
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
