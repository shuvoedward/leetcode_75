package leetcode75

import (
	"fmt"
	"unicode"
)

// not part of 75, 3136
func isVowel(r rune) bool {
	r = unicode.ToLower(r)
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}
func isValidWord(word string) bool {
	if len(word) < 3 {
		return false
	}

	hasVowel := false
	hasConsonant := false

	for _, r := range word {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}

		if unicode.IsLetter(r) {
			if isVowel(r) {
				hasVowel = true
			} else {
				hasConsonant = true
			}
		}
	}

	return hasVowel && hasConsonant

}

func isConsonant(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isVowelManual(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		r += 32 // convert uppercase to lowercase
	}
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
func isValidWordManual(word string) bool {
	if len(word) < 3 {
		return false
	}
	hasVowel, hasConsonant := false, false

	for _, r := range word {
		// check if special character
		if !isConsonant(r) && !isDigit(r) {
			return false
		}

		if isConsonant(r) {
			if isVowelManual(r) {
				hasVowel = true
				fmt.Println("found vowel ", r)
			} else {

				hasConsonant = true
			}
		}
	}
	return hasConsonant && hasVowel
}
