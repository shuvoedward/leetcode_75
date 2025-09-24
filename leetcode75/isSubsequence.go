package leetcode75

func isSubsequence(s, t string) bool {
	i, j := 0, 0

	// focuses on the order of s
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}
