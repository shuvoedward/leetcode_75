package main

func longestOnes(nums []int, k int) int {
	left := 0
	zeroCount := 0
	maxLength := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		currentWindowLength := right - left + 1
		if currentWindowLength > maxLength {
			maxLength = currentWindowLength
		}
	}

	return maxLength
}
