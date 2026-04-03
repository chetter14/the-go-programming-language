package main

import (
	"fmt"
	"popcntsample/popcount"
)

func main() {
	fmt.Println("With table version:")
	fmt.Println(popcount.PopCount(4))
	fmt.Println(popcount.PopCount(8))
	fmt.Println(popcount.PopCount(7))
	fmt.Println(popcount.PopCount(256))
	fmt.Println(popcount.PopCount(255))
	
	fmt.Println("Brute-force version:")
	fmt.Println(popcount.PopCount(4))
	fmt.Println(popcount.PopCount(8))
	fmt.Println(popcount.PopCount(7))
	fmt.Println(popcount.PopCountSlow(7))
	fmt.Println(popcount.PopCountSlow(256))
	fmt.Println(popcount.PopCountSlow(255))
}
