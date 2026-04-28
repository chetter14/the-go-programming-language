package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
}

func foo() (res int) {
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		default:
			res = p.(int)
		}
	}()
	panic(13)
}
