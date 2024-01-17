package main

import (
	"slices"
	"strings"
)

// Fields + Builder
// ----------------------------------------------------------
// func reverseWords(s string) string {
// 	words := strings.Fields(s)

// 	var b strings.Builder
// 	b.Grow(len(s))

// 	for len(words) > 0 {
// 		b.WriteString(" ")
// 		b.WriteString(words[len(words)-1])
// 		words = words[:len(words)-1]
// 	}

// 	return strings.TrimPrefix(b.String(), " ")
// }

// "slices" and "strings"
// ----------------------------------------------------------
func reverseWords(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)

	return strings.Join(words, " ")
}
