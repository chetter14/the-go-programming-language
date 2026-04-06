package main

import (
	"bytes"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))

	fmt.Println(comma("1234567"))

	s := "abc"
	b := []byte(s)
	s2 := string(b)

	fmt.Println(s2)

	fmt.Println(intsToString([]int{1, 2, 3, 4, 5}))

	fmt.Println(commaBuffer("1234567"))

	fmt.Println(areAnagrams("hello", "world"))
	fmt.Println(areAnagrams("x", "world"))
	fmt.Println(areAnagrams("bob", "bob"))
	
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	
	fmt.Println(strconv.FormatInt(int64(x), 2))
	s = fmt.Sprintf("x=%b", x)
	fmt.Println(s)
	
	val1, err := strconv.Atoi("123")
	val2, err := strconv.ParseInt("123", 10, 64)
	
	fmt.Println(val1, val2, err)
}

// Version 1:
// func basename(s string) string {
// for i := len(s) - 1; i >= 0; i-- {
// if s[i] == '/' {
// s = s[i+1:]
// break
// }
// }

// for i := len(s) - 1; i >= 0; i-- {
// if s[i] == '.' {
// s = s[:i]
// break
// }
// }
// return s
// }

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaBuffer(s string) string {
	var buf bytes.Buffer

	commaIndex := len(s) % 3

	for i, c := range s {
		if commaIndex == i {
			fmt.Fprintf(&buf, ",")
			commaIndex += 3
		}
		fmt.Fprintf(&buf, "%c", c)
	}
	return buf.String()
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func areAnagrams(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i1, i2 := 0, len(s2)-1; i1 < (len(s1)/2)+1; i1, i2 = i1+1, i2-1 {
		if s1[i1] != s2[i2] {
			return false
		}
	}
	return true
}
