package arrayhashing

func productExceptSelf(nums []int) []int {
	//  brute force
	// range creates a copy of the array and slightly slower
	// range does not create copy of the whole slice but value of th slice
	// using range not possible to modify the array or slice
	n := len(nums)
	res := make([]int, n)

	for i := 0; i < len(nums); i++ {
		prod := 1
		for m := 0; m <= len(nums); m++ {
			if i != m {
				prod *= nums[m]
			}
		}
		res[i] = prod
		/*
			if i == m {
				continue
			}
		*/

	}
	return res
}

func division(nums []int) []int {
	prod := 1
	zeroCount := 0

	for _, num := range nums {
		if num != 0 {
			prod *= num
		} else {
			zeroCount++
		}
	}

	res := make([]int, len(nums))

	// if only one zero, then its ok. but if more than one then the products will always be zero.
	if zeroCount > 1 {
		return res
	}

	for i, num := range nums {
		if zeroCount > 0 {
			if num == 0 {
				res[i] = prod

			} else {
				res[i] = 0
			}

		} else {
			res[i] = prod / num
		}
	}
	return res

}

func profixAndSuffix(nums []int) []int {
	// Optimal

	res := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		res[i] = prefix
		prefix *= num
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {

		res[i] *= postfix
		postfix *= nums[i]

	}
	return res

}
