package main

import (
	"bufio"
	"io"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"
)

type charCntResult struct {
	counts  map[rune]int
	utflen  [utf8.UTFMax + 1]int
	invalid int
}

func charCount(str string) charCntResult {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(strings.NewReader(str))
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			return charCntResult{nil, utflen, 1}
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	return charCntResult{counts, utflen, invalid}
}

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input string
		res   charCntResult
	}{
		{"", charCntResult{
			map[rune]int{},
			[utf8.UTFMax + 1]int{},
			0}},
		{"abc", charCntResult{
			map[rune]int{'a': 1, 'b': 1, 'c': 1},
			[utf8.UTFMax + 1]int{0, 3},
			0}},
		{"aaAAbbBBccCC", charCntResult{
			map[rune]int{'a': 2, 'b': 2, 'c': 2, 'A': 2, 'B': 2, 'C': 2},
			[utf8.UTFMax + 1]int{0, 12},
			0}},
		{"aAжЖ", charCntResult{
			map[rune]int{'a': 1, 'A': 1, 'ж': 1, 'Ж': 1},
			[utf8.UTFMax + 1]int{0, 2, 2},
			0}},
		{"\xff\xfe", charCntResult{
			nil,
			[utf8.UTFMax + 1]int{},
			2,
		}},
	}

	for _, test := range tests {
		res := charCount(test.input)

		if res.invalid != test.res.invalid {
			t.Errorf("charCount(%q) received invalid = %d, want %d",
				test.input, res.invalid, test.res.invalid)
			continue
		}

		if len(test.res.counts) != len(res.counts) {
			t.Errorf("charCount(%q) received len(counts) = %d, want %d",
				test.input, len(res.counts), len(test.res.counts))
			continue
		}

		for i, cnt := range res.utflen {
			if cnt != test.res.utflen[i] {
				t.Errorf("charCount(%q) received utflen[%d] = %d, want %d",
					test.input, i, cnt, test.res.utflen[i])
				break
			}
		}

		for r, cnt := range res.counts {
			if cnt != test.res.counts[r] {
				t.Errorf("charCount(%q) received counts[%q] = %d, want %d",
					test.input, r, cnt, test.res.counts[r])
				break
			}
		}
	}
}
