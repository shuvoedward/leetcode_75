package leetcode75

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	pre := 1
	for i := 0; i < len(nums); i++ {
		answer[i] = pre
		pre *= nums[i]
	}

	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = answer[i] * post
		post *= nums[i]
	}

	return answer
}

func productExceptSelfBetter(nums []int) []int {
	// [2, 3, 4, 5]

	length := len(nums)
	result := make([]int, length)

	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}
