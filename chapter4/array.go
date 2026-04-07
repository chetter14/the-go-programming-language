package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q)
	fmt.Println(r)

	q2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(q2, len(q2))

	// q2 = [3]int{1, 2, 3}		compile error !

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])
	fmt.Printf("%T\n", symbol)

	/* An array with 100 elements, all zero except for the last, which has value −1 */
	b := [...]int{99: -1}
	fmt.Printf("%T\n", b)

	compareArrays()

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Println(countBitsSha256(c1))
}

func compareArrays() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, b == c, a == c)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
	/* Or */
	*ptr = [32]byte{}
}

func countBits(val byte) int {
	cnt := 0
	for i := 0; i < 8; i++ {
		cnt += (int)((val >> i) & 1)
	}
	return cnt
}

func countBitsSha256(sha256Value [32]byte) int {
	cnt := 0
	for _, v := range sha256Value {
		cnt += countBits(v)
	}
	return cnt
}
