package main

import (
	"fmt"
)

func main() {
	fmt.Println(-5 % -3)
	fmt.Println(-5 % 3)
	fmt.Println(5.0 / 4.0)
	fmt.Println(5 / 4)

	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint8(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	var apples int32 = 1
	var oranges int16 = 2
	var compote int = int(oranges) + int(apples)
	fmt.Println(compote)

	f := 3.141
	fmt.Println(f, int(f))
	f = 1.99
	fmt.Println(int(f))
	f = 1e100
	fmt.Println(int(f))

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	hex := 0xdeadbeef
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", hex)

	ascii := 'a'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
}
