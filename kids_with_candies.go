package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, c := range candies {
		if c > max {
			max = c
		}
	}

	result := make([]bool, len(candies))
	max -= extraCandies

	for i, c := range candies {
		if c >= max {
			result[i] = true
		}
	}

	return result
}
