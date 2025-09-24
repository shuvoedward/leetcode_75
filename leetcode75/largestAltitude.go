package leetcode75

func largestAltitude(gain []int) int {

	highest := 0
	val := 0

	for i := range gain {
		val += gain[i]
		if val > highest {
			highest = val
		}

	}
	return highest
}
