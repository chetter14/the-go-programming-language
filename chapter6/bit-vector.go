package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())    // "3"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println(y.Len())    // "2"

	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	z := x.Copy()
	fmt.Println(x.String()) //  "{1 9 42 144}"
	fmt.Println(z.String()) //  "{1 9 42 144}"

	x.Remove(9)
	fmt.Println(x.String()) //  "{1 42 144}"

	z.Clear()
	fmt.Println(z.String()) //  "{}"

	z.AddAll(1, 42, 123)
	fmt.Println(z.String()) //  "{1 42 123}"
	fmt.Println(z.Len())    // "3"

	arr := []int{9, 42, 123, 144}
	z.Clear()
	z.AddAll(arr...)
	fmt.Println(z.String()) //  "{9, 42, 123, 144}"
	fmt.Println(z.Len())    // "4"

	x.IntersectWith(z)
	fmt.Println(x.String()) // "{42, 144}"

	x.SymmetricDifference(&y)
	fmt.Println(x.String()) // "{9, 144}"

	x.Clear()
	x.AddAll(9, 123, 144)
	x.DifferenceWith(&y)
	fmt.Println(x.String()) // "{123, 144}"

	elems := x.Elems()
	fmt.Printf("%T\n", elems) // "[]int"
	fmt.Println(elems)        // "[123, 144]"

	elems = z.Elems()
	fmt.Println(elems) // "[9, 42, 123, 144]"
}

// An IntSet is a set of small nonnegative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the nonnegative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the nonnegative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Return the number of elements
func (s *IntSet) Len() int {
	bitCount := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				bitCount++
			}
		}
	}
	return bitCount
}

// Remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] = s.words[word] &^ (1 << bit)
	}
}

// Remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Return a copy of the set
func (s *IntSet) Copy() *IntSet {
	return &IntSet{words: append([]uint64(nil), s.words...)}
}

func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// DifferenceWith sets s to the difference of s with t (what's in s but not in t)
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

// SymmetricDifference sets s to the symmetric difference of s and t
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Elems() []int {
	res := make([]int, 0, s.Len())

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, 64*i+j)
			}
		}
	}
	return res
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
