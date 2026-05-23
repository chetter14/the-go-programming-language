package main

import (
	"fmt"
)

type mySortInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func IsPalindrome(s mySortInterface) bool {
	if s.Len() < 2 {
		return true
	}

	i, j := 0, s.Len()-1
	for i < j {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		i++
		j--
	}
	return true
}

type Array []int

func (a Array) Len() int           { return len(a) }
func (a Array) Less(i, j int) bool { return a[i] < a[j] }
func (a Array) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	fmt.Println(IsPalindrome(Array{1, 2, 1}))
	fmt.Println(IsPalindrome(Array{2, 4, 3, 4, 2}))
	fmt.Println(IsPalindrome(Array{1, 2, 3}))
}
