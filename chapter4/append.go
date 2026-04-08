package main

import (
	"fmt"
)

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var a []int
	a = append(a, 1)
	a = append(a, 2, 3)
	a = append(a, 4, 5, 6)
	a = append(a, a...)
	fmt.Println(a)

	var b []int
	b = appendInt(b, 1, 2, 3)
	b = appendInt(b, b...)
	fmt.Println(b)
}

/* 'y...' is a variadic parameter */
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		/* There is room to grow; extend */
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
