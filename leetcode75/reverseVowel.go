package leetcode75

import "strings"

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	isVowel := func(c byte) bool {
		return strings.ContainsRune(vowels, rune(c))
	}

	chars := []byte(s)
	left, right := 0, len(chars)-1

	for left < right {
		for left < right && !isVowel(chars[left]) {
			left++
		}
		for left < right && !isVowel(chars[right]) {
			right--
		}

		if left < right {
			chars[left], chars[right] = chars[right], chars[left]
			left++
			right--
		}
	}
	return string(chars)
}
