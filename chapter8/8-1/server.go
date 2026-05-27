package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	port := flag.String("port", "8000", "port to listen on")
	timezone := flag.String("tz", "UTC", "timezone used to fetch the time")
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+(*port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, *port, *timezone) // handle one connection at a time
	}
}

func handleConn(c net.Conn, port string, timezone string) {
	defer c.Close()

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatal("failed to load a location")
	}

	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
