package main

import (
	"fmt"
	"os"
)

// func main() {
// // Echo the cmd arguments
// var s, sep string
// for i := 1; i < len(os.Args); i++ {
// s += sep + os.Args[i]
// sep = " "

// for i < 0 {
// fmt.Println("It's never going to be printed")
// }
// }
// fmt.Println(s)

// }

// func main() {
// // Short variable declaration
// s, sep := "", ""
// // Unused value in '_'
// for _, arg := range os.Args[1:] {
// s += sep + arg
// sep = " "
// }
// fmt.Println(s)
// }

func main() {
	for index, arg := range os.Args {
		fmt.Print(index)
		fmt.Println(" " + arg)
	}
}

// func main() {
// // fmt.Println(strings.Join(os.Args[1:], " "))
// fmt.Println(os.Args[1:])
// }
