package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 16777216
	fmt.Println(f == f+1)

	const e = 2.71828
	const Avogadro = 6.022

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z, math.IsNaN(z/z), math.IsInf(1/z, 1))

	fmt.Println(compute())
}

func compute() (float64, bool) {
	result, failed := 1+2, false

	if failed {
		return 0, false
	}
	return float64(result), true
}
