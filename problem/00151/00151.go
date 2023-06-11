package _00151

import (
	"strings"
)

// 151. Reverse Words in a String
// https://leetcode.com/problems/reverse-words-in-a-string

func reverseWords(s string) string {
	tmp, ans := strings.Split(s, " "), ""
	for i := len(tmp) - 1; i >= 0; i-- {
		if tmp[i] == "" {
			continue
		}
		ans += " " + tmp[i]
	}
	return ans[1:]
}
