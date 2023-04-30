package _00003

import (
	"math"
)

// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	m := map[byte]float64{}
	ans := 0.0

	for i, c := 0, 1.0; i < len(s); i, c = i+1, c+1 {
		if _, ok := m[s[i]]; ok {
			cc := m[s[i]]
			c -= cc
			for j := range m {
				m[j] -= cc
				if m[j] <= 0 {
					delete(m, j)
				}
			}
		}
		m[s[i]] = c
		ans = math.Max(c, ans)
	}
	return int(ans)
}
