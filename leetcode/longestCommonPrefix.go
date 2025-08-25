package leetcode

func LongestCommonPrefix(strs []string) string {
	smallestWord := strs[0]
	result := ""

	for _, word := range strs {
		if len(smallestWord) > len(word) {
			smallestWord = word
		}
	}

	for i, _ := range smallestWord {
		match := false
		for _, word := range strs {
			if smallestWord[i] == word[i] {
				match = true
			} else {
				match = false
			}
		}
		if match {
			result += string(smallestWord[i])
		} else {
			return result
		}
	}

	return result
}

func LongestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	firstString := strs[0]

	for i, char := range firstString {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != byte(char) {
				return firstString[:i]
			}
		}
	}
	return firstString
}
