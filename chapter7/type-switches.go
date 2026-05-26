package main

import (
	"fmt"
	"os"
)

func getType(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return "Integer (signed or unsigned)"
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return "String"
	default:
		return fmt.Sprintf("Unexpected type %T: %v", x, x)
	}
}

func main() {
	fmt.Println(getType(2))
	fmt.Println(getType(uint(3)))
	fmt.Println(getType(true))
	fmt.Println(getType(nil))
	fmt.Println(getType("text"))
	fmt.Println(getType(os.File{}))
	fmt.Println("Completed")
}
