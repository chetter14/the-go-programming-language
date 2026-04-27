package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(1))
	fmt.Println(max(1, 2))
	fmt.Println(max(1, 2, 3))
	fmt.Println(max(4, 1, 8, 6))

	fmt.Println(min(1))
	fmt.Println(min(3, 2))
	fmt.Println(min(7, 6, 5))
}

func max(val int, vals ...int) int {
	res := val
	for _, x := range vals {
		if x > res {
			res = x
		}
	}
	return res
}

func min(val int, vals ...int) int {
	res := val
	for _, x := range vals {
		if x < res {
			res = x
		}
	}
	return res
}
