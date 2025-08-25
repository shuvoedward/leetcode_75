package main

import (
	"strings"
)

func reverseWords(s string) string {
	// the sky is blue
	words := strings.Fields(s)

	// reverse the words slice
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
