package main

func increasingTriplet(nums []int) bool {
	// 2, 5, 4, 0, 6
	if len(nums) < 3 {
		return false
	}

	INF := int(^uint(0) >> 1)
	first, second := INF, INF

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			// x > second happens, shows that, first and second already found

			return true
		}

	}
	return false
}
