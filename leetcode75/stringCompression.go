package leetcode75

import "fmt"

func Compress(chars []byte) int {
	n := len(chars)
	write := 0 // Position to write compressed character
	read := 0  // position to read characters

	for read < n {
		char := chars[read]
		count := 0

		// count the numbers of repeated characters
		for read < n && chars[read] == char {
			read++
			count++
		}

		// Write the characters
		chars[write] = char
		write++

		// If count > 1, write the count as characters
		if count > 1 {
			countStr := fmt.Appendf(nil, "%d", count)
			for _, digit := range countStr {
				chars[write] = digit
				write++
			}
		}

	}
	return write
}
