package main

func largestAltitude(gain []int) int {
	var max, altitude int

	for _, g := range gain {
		altitude += g
		if altitude > max {
			max = altitude
		}
	}

	return max
}
