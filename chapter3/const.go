package main

import (
	"fmt"
	"time"
)

const (
	e  = 2.71
	pi = 3.14
)

const (
	a = 1
	b
	c = 2
	d
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(e, pi)

	const IPv4Len = 4

	var p [IPv4Len]byte
	p[0] = 0xFF

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	fmt.Println(a, b, c, d)
	fmt.Println(Sunday, Monday, Thursday, Saturday)

	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)
	fmt.Println(5 / 9 * (f - 32))
	fmt.Println(5.0 / 9.0 * (f - 32))

	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000')

	var j = int8(0)
	fmt.Printf("%T\n", j)
}
