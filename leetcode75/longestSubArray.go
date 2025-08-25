package main

func longestSubarray(nums []int) int {
	left := 0
	zeroCount := 0
	maxLength := 0

	for right, num := range nums {
		if num == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		currentWindowLength := right - left
		if currentWindowLength > maxLength {
			maxLength = currentWindowLength
		}
	}
	return maxLength
}
