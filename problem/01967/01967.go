package _01967

import "strings"

// 1967. Number of Strings That Appear as Substrings in Word
// https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/

func numOfStrings(patterns []string, word string) int {
	ans := 0
	for _, s := range patterns {
		if strings.Contains(word, s) {
			ans++
		}
	}
	return ans
}
