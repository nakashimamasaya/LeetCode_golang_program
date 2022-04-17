package _0028

import "strings"

// 28. Implement strStr()
// https://leetcode.com/problems/implement-strstr/
func StrStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
