package main

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, piles[0]

	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}

	res := right
	for left <= right {
		mid := left + (right-left)/2
		totalTime := 0

		for _, banana := range piles {
			totalTime += (banana + mid - 1) / mid
		}

		if totalTime <= h {
			right = mid - 1
			res = mid
		} else {
			left = mid + 1
		}
	}
	return res
}
