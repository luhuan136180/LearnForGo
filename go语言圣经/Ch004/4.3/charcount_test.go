package main

import (
	"bytes"
	"testing"
)

func TestCharCount(t *testing.T) {
	str := &bytes.Buffer{}
	str.WriteString("hello 世界")
	counts, utflen, invalid := charcount(str)

	countsWant := map[rune]int{
		'h': 1,
		'e': 1,
		'l': 2,
		'o': 1,
		'世': 1,
		'界': 1,
		' ': 1,
	}
	utflenWant := []int{0, 6, 0, 2, 0}
	invalidWant := 0

	if invalid != invalidWant {
		t.Errorf("invalid count = %v, want %v", invalid, invalidWant)
	}

	for c, n := range countsWant {
		if got, ok := counts[c]; !ok || got != n {
			t.Errorf("counts = %v, want %v", counts, countsWant)
		}
	}

	if len(utflenWant) != len(utflen) {
		t.Errorf("utflen = %v, want %v", utflen, utflenWant)
	}

	for i := range utflenWant {
		if utflen[i] != utflenWant[i] {
			t.Errorf("utflen = %v, want %v", utflen, utflenWant)
			break
		}
	}
}
