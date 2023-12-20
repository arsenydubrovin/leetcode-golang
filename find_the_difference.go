package main

// Sums of codes
// ----------------------------------------------------------
// func findTheDifference(s string, t string) byte {
// 	var sCodeSum, tCodeSum int
// 	for _, l := range s {
// 		sCodeSum += int(l)
// 	}

// 	for _, l := range t {
// 		tCodeSum += int(l)
// 	}

// 	return byte(tCodeSum - sCodeSum)
// }

// Increment and decrement (Not my)
// ----------------------------------------------------------
func findTheDifference(s string, t string) byte {
	var result int
	for _, l := range s {
		result += int(l)
	}

	for _, l := range t {
		result -= int(l)
	}

	return byte(result)
}
