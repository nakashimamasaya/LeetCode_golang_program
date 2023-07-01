package _00392

// 392. Is Subsequence
// https://leetcode.com/problems/is-subsequence

func isSubsequence(s string, t string) bool {
	i := 0
	for j := 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}
