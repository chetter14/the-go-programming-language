package main

import "testing"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	total_cnt := 0
	for i := 0; i < 8; i++ {
		total_cnt += int(pc[byte(x>>(i*8))])
	}
	return total_cnt
}

func PopCountSlow(x uint64) int {
	total_cnt := 0
	for i := 0; i < 64; i++ {
		total_cnt += int((x >> i) & 1)
	}
	return total_cnt
}

var sink int

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = PopCount(uint64(i))
	}
}

func BenchmarkPopCountSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = PopCountSlow(uint64(i))
	}
}
