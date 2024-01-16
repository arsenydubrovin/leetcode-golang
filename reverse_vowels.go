package main

func reverseVowels(s string) string {
	var stack []rune

	for _, letter := range s {
		if isVowel(letter) {
			stack = append(stack, letter)
		}
	}

	var answer []rune
	for _, letter := range s {
		if isVowel(letter) {
			letter, stack = stack[len(stack)-1], stack[:len(stack)-1]
		}
		answer = append(answer, letter)
	}

	return string(answer)
}

func isVowel(letter rune) bool {
	for _, vowel := range []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'} {
		if letter == vowel {
			return true
		}
	}
	return false
}

// Better isVowel (not my)
// ----------------------------------------------------------
// func isVowel(r rune) bool {
//     switch r {
//         case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
//             return true
//         default:
//             return false
//     }
// }

// Optimized memory with strings.Builder
// ----------------------------------------------------------
// func reverseVowels(s string) string {
// 	var stack []rune

// 	for _, letter := range s {
// 		if isVowel(letter) {
// 			stack = append(stack, letter)
// 		}
// 	}

// 	var answer strings.Builder
// 	answer.Grow(len(s))

// 	for _, letter := range s {
// 		if isVowel(letter) {
// 			letter, stack = stack[len(stack)-1], stack[:len(stack)-1]
// 		}
// 		answer.WriteRune(letter)
// 	}

// 	return answer.String()
// }
