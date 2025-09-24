package leetcode

func IsValidParantheses(s string) bool {
	stack := []rune{}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if openBracket, isClosingBracket := pairs[char]; isClosingBracket {
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != openBracket {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
