package leetcode75

import "strings"

func mergeStrings(word1, word2 string) string {
	result := ""

	len1 := len(word1)
	len2 := len(word2)

	maxLen := max(len1, len2)

	for i := range maxLen {
		if i < len1 {
			result += string(word1[i])
		}
		if i < len2 {
			result += string(word2[i])
		}
	}
	return result

	// concatenation creates a new memory for the string each time.
}

// Using strings builder
func mergeStrings2(word1, word2 string) string {
	var result strings.Builder

	len1 := len(word1)
	len2 := len(word2)

	maxLen := len1
	if len2 > len1 {
		maxLen = len2
	}

	for i := range maxLen {
		if i < len1 {
			result.WriteByte(word1[i])
		}
		if i < len2 {
			result.WriteByte(word2[i])
		}
	}

	return result.String()
}

func mergeStrings3(word1, word2 string) string {
	var result strings.Builder

	minLen := min(len(word2), len(word1))

	for i := range minLen {
		result.WriteByte(word1[i])
		result.WriteByte(word2[i])
	}

	if len(word1) > minLen {
		result.WriteString(word1[minLen:])
	}
	if len(word2) > minLen {
		result.WriteString(word2[minLen:])
	}

	return result.String()
}
