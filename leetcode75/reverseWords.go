package leetcode75

import (
	"strings"
	"unicode"
)

func ReverseWords(s string) string {
	// the sky is blue
	words := strings.Fields(s)

	// reverse the words slice
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
func ReverseWord2(s string) string {
	words := []string{}
	wordstart := 0

	s = strings.TrimSpace(s)

	for i, r := range s {
		if unicode.IsSpace(r) {
			word := s[wordstart:i]
			if len(word) > 0 {
				words = append(words, word)
			}
			wordstart = i + 1
		}
	}

	lastWord := s[wordstart:]
	words = append(words, lastWord)

	var result strings.Builder
	for i := len(words) - 1; i >= 0; i-- {
		result.WriteString(words[i])
		if i > 0 {
			result.WriteString(" ")
		}
	}

	return result.String()
}
