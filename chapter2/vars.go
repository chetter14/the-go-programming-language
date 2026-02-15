package main

import "fmt"

func main() {
	var s string
	fmt.Println(s)

	var b, f, t = false, 4.2, "text"
	fmt.Printf("%t %g %s\n", b, f, t)

	/* Short variable declaration */
	i := 100
	/* Assignment */
	i = 10

	j := 50
	i, j = j, i /* Swap */

	fmt.Printf("i - %d, j - %d\n", i, j)

	k, i := 30, 187
	fmt.Printf("i - %d, j - %d, k - %d\n", i, j, k)

	// j, k := 42, 17 	/* compile error */

	x := 1
	p := &x
	fmt.Printf("x contains - %d\n", *p) // 1
	*p = 2
	fmt.Printf("x contains - %d\n", *p) // 2

	fmt.Println(foo() == foo())

	x = 10
	inc(&x)
	fmt.Printf("x contains - %d\n", x)
}

func foo() *int {
	var x = 1
	return &x
}

func inc(p *int) int {
	*p++
	return *p
}
