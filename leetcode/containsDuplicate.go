package leetcode

func ContainsDuplicate(nums []int) bool {
	count := make(map[int]bool)

	for _, num := range nums {
		if _, exist := count[num]; exist {
			return true
		}
		count[num] = true
	}
	return false
}
