package main

import (
	"fmt"
)

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// fmt.Println(summer[:20])	// panic !

	/* Extend the slice (within capacity) */
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)

	/* Array */
	a := [...]int{1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	/* Slice (array is inited implicitly) */
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)

	fmt.Println(equal([]string{"hello", "world"}, []string{"world", "hello"}))
	fmt.Println(equal([]string{"hello", "world"}, []string{"hello", "world"}))
	
	var sl []int
	sl = nil
	sl = []int(nil)
	fmt.Println(sl == nil)
	sl = []int{}
	fmt.Println(sl == nil)
	fmt.Printf("%T %d\n", sl, len(sl))
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
