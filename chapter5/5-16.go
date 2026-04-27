package main

import (
	"fmt"
)

func main() {
	fmt.Println(stringsJoin(" ", "hello", "world"))

	words := []string{"My", "name", "is"}
	fmt.Println(stringsJoin(" ", words...))
}

func stringsJoin(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, str := range strs[1:] {
		res += sep + str
	}
	return res
}
