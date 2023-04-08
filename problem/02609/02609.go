package _02609

import "math"

// 2609. Find the Longest Balanced Substring of a Binary String
// https://leetcode.com/problems/find-the-longest-balanced-substring-of-a-binary-string/

func findTheLongestBalancedSubstring(s string) int {
	ans := 0.0
	for i := 0; i < len(s); i++ {
		if s[i] == 48 {
			count, c := 0, 0.0
			for ; i < len(s) && s[i] == 48; i, count = i+1, count+1 {
			}
			l := int(math.Min(float64(i+count), float64(len(s))))
			for ; i < l && s[i] == 49; i, c = i+1, c+1 {
			}
			ans, i = math.Max(ans, c*2), i-1
		}
	}
	return int(ans)
}
