package leetcode75

import "sort"

func successfulPairsBruteForce(spells []int, potions []int, success int64) []int {
	result := make([]int, len(spells))

	for i, spell := range spells {
		for _, potion := range potions {
			if int64(spell*potion) >= success {
				result[i]++
			}
		}

	}
	return result
}

func successfulPairsBinary(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	n := len(spells)
	m := len(potions)
	pairs := make([]int, n)

	for i, spell := range spells {
		var minPotionStrength int

		if success%int64(spell) == 0 {
			minPotionStrength = int(success / int64(spell))
		} else {
			minPotionStrength = int(success/int64(spell)) + 1
		}

		idx := sort.SearchInts(potions, minPotionStrength)

		pairs[i] = m - idx
	}

	return pairs

	// time complexity nlogm
}

func successfulPairsPrefixSum(spells []int, potions []int, success int64) []int {
	suc := int(success)
	maxp := 0
	for _, p := range potions {
		maxp = max(maxp, p)
	}

	dp := make([]int, maxp+1)
	for _, p := range potions {
		dp[p]++
	}

	// total number of potions that have a strength of i or greater
	for i := maxp - 1; i >= 0; i-- {
		dp[i] += dp[i+1]
	}

	res := make([]int, len(spells))
	for i, s := range spells {
		// s * p >= suc; p >= suc / s
		need := (suc + s - 1) / s
		if need <= maxp {
			res[i] = dp[need]
		}
	}
	return res
}
