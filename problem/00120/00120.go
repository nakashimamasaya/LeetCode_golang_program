package _00120

import (
	"sort"
)

// 120. Triangle
// https://leetcode.com/problems/triangle/

func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] += triangle[i-1][0]
		triangle[i][i] += triangle[i-1][i-1]
		for j := 1; j < i; j++ {
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
	}
	ans := triangle[len(triangle)-1]
	sort.Ints(ans)
	return ans[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
