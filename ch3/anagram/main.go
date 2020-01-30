package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

//kürek ürkek true
//ev ve true
//adam madam false

var w1 = []string{"kürek", "ev", "adam"}
var w2 = []string{"ürkek", "ve", "madam"}

func main() {
	for i, w := range w1 {
		fmt.Printf("Byte version: w1:%s w2:%s res:%t\n", w, w2[i], isAnagramByte(w, w2[i]))
		fmt.Printf("Rune version: w1:%s w2:%s res:%t\n", w, w2[i], isAnagramRune(w, w2[i]))
		fmt.Printf("String version: w1:%s w2:%s res:%t\n", w, w2[i], isAnagramRune(w, w2[i]))
	}
}

func isAnagramByte(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	b1 := []byte(s1)
	b2 := []byte(s2)

	sort.Slice(b1, func(i int, j int) bool { return b1[i] < b1[j] })
	sort.Slice(b2, func(i int, j int) bool { return b2[i] < b2[j] })

	for i, b := range b1 {
		if b2[i] != b {
			return false
		}
	}

	return true
}

func isAnagramRune(s1, s2 string) bool {
	if utf8.RuneCountInString(s1) != utf8.RuneCountInString(s2) {
		return false
	}

	r1 := []rune(s1)
	r2 := []rune(s2)

	sort.Slice(r1, func(i int, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i int, j int) bool { return r2[i] < r2[j] })

	for i, b := range r1 {
		if r2[i] != b {
			return false
		}
	}

	return true
}

func isAnagramString(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	ss1 := strings.Split(s1, "")
	ss2 := strings.Split(s2, "")

	sort.Strings(ss1)
	sort.Strings(ss2)

	for i, b := range ss1 {
		if ss2[i] != b {
			return false
		}
	}

	return true
}
