package _00038

import "strconv"

// 38. Count and Say
// https://leetcode.com/problems/count-and-say/

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s, ans, count := countAndSay(n-1), "", 1

	for i := 1; i < len(s); i, count = i+1, count+1 {
		if s[i-1] != s[i] {
			ans += strconv.Itoa(count) + string(s[i-1])
			count = 0
		}
	}
	return ans + strconv.Itoa(count) + string(s[len(s)-1])
}
