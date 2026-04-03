package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func f() int { return 1 }

func g(val int) int { return 1 }

// func main() {
// f := "f"
// fmt.Println(f)
// fmt.Println(g)
// // fmt.Println(h)	// compile error !
// }

// func main() {
// x := "hello"
// for _, x := range x {
// x := x + 'A' - 'a'
// fmt.Printf("%c", x)
// }
// }

func main() {
	if x := f(); x == 0 {
		fmt.Println("If block")
		fmt.Println(x)
	} else if y := g(x); y == x {
		fmt.Println("Else-if block")
		fmt.Println(x, y)
	} else {
		fmt.Println("Else block")
		fmt.Println(x, y)
	}

	// fmt.Println(x, y)	// compile error !

	f, err := os.Open("test_file")
	if err != nil {
		fmt.Println("Failed to open")
		return
	}

	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Failed to read", err)
		f.Close()
		return
	}

	fmt.Println(string(content))
	f.Close()
}

var cwd string

func init() {
	// cwd, err := os.Getwd()	// WRONG! 'cwd' is a local variable here, global 'cwd' is not assigned any value
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}
