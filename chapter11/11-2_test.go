package main

import (
	"bytes"
	"fmt"
	"testing"
)

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

func TestIntSetHas(t *testing.T) {
	x := IntSet{[]uint64{7, 1}}

	var tests = []struct {
		input int
		want  bool
	}{
		{0, true},
		{1, true},
		{2, true},
		{3, false},
		{32, false},
		{63, false},
		{64, true},
		{65, false},
	}

	for _, test := range tests {
		if got := x.Has(test.input); got != test.want {
			t.Errorf("IntSet.Has(%d) = %v", test.input, got)
		}
	}
}

func TestIntSetAdd(t *testing.T) {
	x := IntSet{[]uint64{0}}

	var tests = []int{0, 1, 2, 8, 63, 64, 256}

	for _, test := range tests {
		x.Add(test)
		if got := x.Has(test); got != true {
			t.Errorf("IntSet.Has(%d) = false (after adding)", test)
		}
	}
}

func TestIntSetString(t *testing.T) {
	var tests = []struct {
		input IntSet
		res   string
	}{
		{IntSet{[]uint64{}}, "{}"},
		{IntSet{[]uint64{7}}, "{0 1 2}"},
		{IntSet{[]uint64{8}}, "{3}"},
		{IntSet{[]uint64{32}}, "{5}"},
		{IntSet{[]uint64{63}}, "{0 1 2 3 4 5}"},
		{IntSet{[]uint64{255}}, "{0 1 2 3 4 5 6 7}"},
		{IntSet{[]uint64{256}}, "{8}"},
		{IntSet{[]uint64{257}}, "{0 8}"},
	}

	for _, test := range tests {
		if got := test.input.String(); got != test.res {
			t.Errorf("IntSet.String() got = %s, expect = %s", got, test.res)
		}
	}
}
