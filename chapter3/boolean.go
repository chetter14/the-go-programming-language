package main

import (
	"fmt"
)

func main() {
	test("")
	test("xy")

	fmt.Println(btoi(true))
	fmt.Println(btoi(false))

	fmt.Println(itob(17))
	fmt.Println(itob(-1))
	fmt.Println(itob(0))
}

func test(s string) {
	if s != "" && s[0] == 'x' {
		fmt.Println("In if-branch")
	} else {
		fmt.Println("In else-branch")
	}
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0
}
