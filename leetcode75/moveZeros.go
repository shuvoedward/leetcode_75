package leetcode75

import "fmt"

func MoveZeros(nums []int) {
	i := len(nums) - 1
	for i >= 0 {
		if nums[i] != 0 {
			i--
		} else {
			j := i + 1
			k := i
			for j < len(nums) {
				nums[k], nums[j] = nums[k], nums[i]
				j++
				k++
			}
		}
	}
	fmt.Println(nums)
}
