package main

func findPeakElement(nums []int) int {
	var isPeak bool

	for i, num := range nums {
		if i == 0 {
			isPeak = num > nums[i+1]
		} else if i == len(nums)-1 {
			isPeak = num > nums[i-1]
		} else {
			isPeak = num > nums[i+1] && num > nums[i-1]
		}
		if isPeak == true {
			return i
		}
	}

	return 0
}

func findPeakOptimizedBrute(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		leftNeighbor := -1 << 63
		if i > 0 {
			leftNeighbor = nums[i-1]
		}

		rightNeighbor := -1 << 63
		if i < n-1 {
			rightNeighbor = nums[i+1]
		}

		if nums[i] > leftNeighbor && nums[i] > rightNeighbor {
			return i
		}

	}
	return 0
}

func findPeakElementBinarySearch(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
