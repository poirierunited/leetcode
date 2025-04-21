package main

import (
	"fmt"
)

func main() {
	s := "cbbd"

	charMap := make(map[int]string)

	for i, char := range s {
		runeToChar := string(char)
		charMap[i] = runeToChar
	}

	result := longestPalindrome(s)
	fmt.Println(result)
}

// CORRECT APPROACH
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLength := 0, 1

	for i := 0; i < len(s); {
		// Skip duplicate characters (optimization)
		left, right := i, i
		for right < len(s)-1 && s[right] == s[right+1] {
			right++
		}
		i = right + 1

		// Expand around center
		for left > 0 && right < len(s)-1 && s[left-1] == s[right+1] {
			left--
			right++
		}

		// Update max length if needed
		if newLength := right - left + 1; newLength > maxLength {
			start = left
			maxLength = newLength
		}
	}

	return s[start : start+maxLength]
}
