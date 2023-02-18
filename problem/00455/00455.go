package _00455

import "sort"

// 455. Assign Cookies
// https://leetcode.com/problems/assign-cookies/

func findContentChildren(g []int, s []int) int {
	ans := 0

	sort.Ints(g)
	sort.Ints(s)
	for i, j := 0, 0; i < len(g) && j < len(s); j++ {
		if g[i] <= s[j] {
			i, ans = i+1, ans+1
		}
	}
	return ans
}
