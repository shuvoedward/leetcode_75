package main

func findDifference(nums1, nums2 []int) [][]int {
	result := make([][]int, 2)

	m := make(map[int]bool)

	for _, value := range nums1 {
		m[value] = true
	}

	for _, v2 := range nums2 {
		if _, ok := m[v2]; !ok {
			result[1] = append(result[1], v2)
		}
		m[v2] = false
	}

	for key, t := range m {
		if t {
			result[0] = append(result[0], key)
		}
	}

	return result

}
