package _00290

import (
	"strings"
)

// 290. Word Pattern
// https://leetcode.com/problems/word-pattern/

func wordPattern(pattern string, s string) bool {
	m, word := map[byte]string{}, map[string]bool{}

	sSplit := strings.Split(s, " ")

	if len(pattern) != len(sSplit) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if val, ok := m[pattern[i]]; !ok {
			if _, o := word[sSplit[i]]; o {
				return false
			}
			word[sSplit[i]] = true
			m[pattern[i]] = sSplit[i]
		} else if val != sSplit[i] {
			return false
		}
	}
	return true
}
