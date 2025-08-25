package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		if height[left] < height[right] {
			maxArea = max(maxArea, height[left]*(right-left))
			left++
		} else {
			maxArea = max(maxArea, height[right]*(right-left))
			right--
		}
	}
	return maxArea

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
