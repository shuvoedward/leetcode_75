package leetcode75

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	leftSum := 0

	for i := 0; i < len(nums); i++ {
		rightSum := total - leftSum - nums[i]

		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}
