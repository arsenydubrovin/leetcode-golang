package main

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	var l int
	if len(str1) > len(str2) {
		l = len(str2)
	} else {
		l = len(str1)
	}

	for i := l; i >= 1; i-- {
		if (len(str1)%i != 0) || (len(str2)%i != 0) {
			continue
		}
		substr := str1[:i]

		if (strings.Repeat(substr, len(str1)/i) == str1) &&
			(strings.Repeat(substr, len(str2)/i) == str2) {
			return substr
		}
	}
	return ""
}

// First find the real gcd
// ----------------------------------------------------------
// func gcdOfStrings(str1 string, str2 string) string {
// 	a, b := len(str1), len(str2)

// 	for b != 0 {
// 		a, b = b, a%b
// 	}
// 	gcdLength := a

// 	gcdStr := str1[:gcdLength]

// 	if isDivisor(str1, gcdStr) && isDivisor(str2, gcdStr) {
// 		return gcdStr
// 	}

// 	return ""
// }

// func isDivisor(s, divisor string) bool {
// 	n := len(s)
// 	m := len(divisor)

// 	if n%m != 0 {
// 		return false
// 	}

// 	return s == strings.Repeat(divisor, n/m)
// }
