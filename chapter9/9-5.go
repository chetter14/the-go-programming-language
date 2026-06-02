package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	ping := make(chan string)
	pong := make(chan string)
	msg := "data"
	var cnt int64

	go func() {
		pong <- msg

		for range ping {
			atomic.AddInt64(&cnt, 1)
			pong <- msg
		}
	}()

	go func() {
		for range pong {
			atomic.AddInt64(&cnt, 1)
			ping <- msg
		}
	}()

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		n := atomic.SwapInt64(&cnt, 0)
		fmt.Printf("Count per second - %d\n", n)
	}
}
