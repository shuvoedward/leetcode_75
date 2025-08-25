package main

func maxOperations(nums []int, k int) int {
	count := make(map[int]int)
	operations := 0

	for _, num := range nums {
		compliment := k - num
		if count[compliment] > 0 {
			operations++
			count[compliment]--
		} else {
			count[num]++
		}

	}

	return operations
}
