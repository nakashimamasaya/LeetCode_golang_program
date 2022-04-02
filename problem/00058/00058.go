package _0058

import (
	"strings"
)

// 58. Length of Last Word
// https://leetcode.com/problems/length-of-last-word/

func LengthOfLastWord(s string) int {
	sSlice := strings.Split(s, " ")
	for i := len(sSlice) - 1; i >= 0; i-- {
		if len(sSlice[i]) > 0 {
			return len(sSlice[i])
		}
	}
	return 0
}
