package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, c := range `\|/-` {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFib(%d) - %d\n", n, fibN)
}
