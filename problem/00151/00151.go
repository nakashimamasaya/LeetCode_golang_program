package _00151

import (
	"strings"
)

// 151. Reverse Words in a String
// https://leetcode.com/problems/reverse-words-in-a-string/

func reverseWords(s string) string {
	return strings.Join(reverseSlice(strings.Fields(s)), " ")
}

func reverseSlice(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}
