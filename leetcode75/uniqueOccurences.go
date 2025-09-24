package leetcode75

func uniqueOccurences(arr []int) bool {

	frequencyMap := make(map[int]int)
	for _, num := range arr {
		frequencyMap[num]++
	}

	occurenceCount := make(map[int]bool)
	for _, count := range frequencyMap {
		if occurenceCount[count] {
			return false
		}
		occurenceCount[count] = true

	}
	return true

}
