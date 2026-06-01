package main

import (
	"fmt"
	"sync"
)

var pc [256]byte

func fillPopCount() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

var fillPopCountOnce sync.Once

func PopCount(x uint64) int {
	fillPopCountOnce.Do(fillPopCount)
	total_cnt := 0
	for i := 0; i < 8; i++ {
		total_cnt += int(pc[byte(x>>(i*8))])
	}
	return total_cnt
}

func main() {
	fmt.Println(PopCount(4))
	fmt.Println(PopCount(8))
	fmt.Println(PopCount(7))
	fmt.Println(PopCount(256))
	fmt.Println(PopCount(255))
}
