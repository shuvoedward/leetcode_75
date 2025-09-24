package leetcode

func NumJewelsInStones(jewel string, stones string) int {
	count := 0

	for _, j := range jewel {
		for _, s := range stones {
			if j == s {
				count++
			}
		}
	}
	return count
}

func JewelCountHashMap(jewels, stones string) int {
	result := 0
	countJewel := make(map[rune]int)

	for _, jewel := range jewels {
		countJewel[jewel]++
	}

	for _, stone := range stones {
		if count, exist := countJewel[stone]; exist {
			result += count
		}

		// result += countJewel[stone]
	}
	return result
}
