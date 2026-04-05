package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	// c := s[len(s)] 		// panic !

	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	// Errors:
	// fmt.Println(s[-1:])
	// fmt.Println(s[2:1])

	a := "left foot"
	b := a
	a += ", right foot"

	fmt.Println(a)
	fmt.Println(b)

	// a[0] = 'L'		// error !

	a = "data\t'data'"
	b = `data\t'data'`

	fmt.Println(a)
	fmt.Println(b)

	a = "text data here"
	fmt.Println(HasPrefix(a, "text"))
	fmt.Println(HasSuffix(a, "here"))
	fmt.Println(Contains(a, "data"))

	s = "Hello, 汉字"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for range s {
		n++
	}
	fmt.Println(n)

	s = "コンピュータ"
	fmt.Printf("% x\n", s)
	// Returns a sequence of Unicode code points
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))

	// Interprets each value as a rune -> converts using utf-8 encoding
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Println(string(1234567))
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
