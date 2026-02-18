package main

import "fmt"

func main() {
	v := 1
	v++
	v--

	fmt.Println(gcd(10, 15))
	fmt.Println(gcd(21, 6))

	fmt.Println(fib(7))
	fmt.Println(fib(9))

	medals := []string{"gold", "silver", "bronze"}
}

func gcd(x, y int) int {
	for y != 0 {
		/* Tuple assignment */
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		/* Tuple assignment */
		x, y = y, x+y
	}
	return x
}
