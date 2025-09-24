package leetcode

func ContainsDuplicate2(nums []int, k int) bool {
	count := map[int]int{}

	for i, num := range nums {
		if index, exists := count[num]; exists && k >= i-index {
			return true
		}
		count[num] = i
	}
	return false
}
