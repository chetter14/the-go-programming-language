package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s         string
		sep       string
		wantedLen int
	}{
		{"a:b:c", ":", 3},
		{"a, b, c", ",", 3},
		{"it is a sentence", " ", 4},
		{"", "", 0},
		{"some data is here", "", 17},
		{"", " ", 1},
	}

	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.wantedLen {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.s, test.sep, got, test.wantedLen)
		}
	}
}
