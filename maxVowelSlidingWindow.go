package main

func maxVowels(s string, k int) int {

	if k > len(s) || k < 1 {
		return 0
	}

	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	count, maxCount := 0, 0

	// Initialize the first window
	for i := 0; i < k && i < len(s); i++ {
		if isVowel(s[i]) {
			count++
		}
	}

	maxCount = count

	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			count--
		}
		if isVowel(s[i]) {
			count++
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
