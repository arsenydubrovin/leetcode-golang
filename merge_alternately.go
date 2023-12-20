package main

import "strings"

// strings.Builder
// ----------------------------------------------------------
// func mergeAlternately(word1 string, word2 string) string {
// 	var (
// 		i, j    int
// 		builder strings.Builder
// 	)
// 	for {
// 		if i < len(word1) {
// 			builder.WriteRune(rune(word1[i]))
// 			i++
// 		}
// 		if j < len(word2) {
// 			builder.WriteRune(rune(word2[j]))
// 			j++
// 		}

// 		if j == len(word2) && i == len(word1) {
// 			return builder.String()
// 		}
// 	}
// }

// Slice of byte
// ----------------------------------------------------------
// func mergeAlternately(word1 string, word2 string) string {
// 	var i, j int

// 	result := make([]byte, 0, len(word1)+len(word2))

// 	for {
// 		if i < len(word1) {
// 			result = append(result, word1[i])
// 			i++
// 		}
// 		if j < len(word2) {
// 			result = append(result, word2[j])
// 			j++
// 		}
// 		if j == len(word2) && i == len(word1) {
// 			return string(result)
// 		}
// 	}
// }

// Min len and slice of byte
// ----------------------------------------------------------
// func mergeAlternately(word1 string, word2 string) string {
// 	result := make([]byte, 0, len(word1)+len(word2))
// 	minLength := len(word1)

// 	if len(word2) < len(word1) {
// 		minLength = len(word2)
// 	}

// 	for i := 0; i < minLength; i++ {
// 		result = append(result, word1[i], word2[i])
// 	}

// 	result = append(result, word1[minLength:]...)
// 	result = append(result, word2[minLength:]...)

// 	return string(result)
// }

// strings.Builder and one pointer (Not my)
// ----------------------------------------------------------
func mergeAlternately(word1 string, word2 string) string {
	var builder strings.Builder
	builder.Grow(len(word1) + len(word2))

	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			builder.WriteByte(word1[i])
		}
		if i < len(word2) {
			builder.WriteByte(word2[i])
		}
	}
	return builder.String()
}
