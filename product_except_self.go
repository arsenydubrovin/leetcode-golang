package main

// TODO: increase speed!
func productExceptSelf(nums []int) []int {
	leftProducts := make([]int, len(nums))
	rightProducts := make([]int, len(nums))

	leftProducts[0] = 1
	for i := 1; i < len(nums); i++ {
		leftProducts[i] = leftProducts[i-1] * nums[i-1]
	}

	rightProducts[len(rightProducts)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightProducts[i] = rightProducts[i+1] * nums[i+1]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = leftProducts[i] * rightProducts[i]
	}

	return result
}
