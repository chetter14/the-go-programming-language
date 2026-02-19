package main

import (
	"fmt"
	"popcntsample/popcount"
)

func main() {
	fmt.Println(popcount.PopCount(4))
	fmt.Println(popcount.PopCount(8))
	fmt.Println(popcount.PopCount(7))
}
