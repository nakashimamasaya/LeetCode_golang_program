package _02500

import (
	"math"
	"sort"
)

// 2500. Delete Greatest Value in Each Row
// https://leetcode.com/problems/delete-greatest-value-in-each-row/

func deleteGreatestValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}

	ans := 0

	for i := 0; i < len(grid[0]); i++ {
		tmp := grid[0][i]
		for j := 1; j < len(grid); j++ {
			tmp = int(math.Max(float64(tmp), float64(grid[j][i])))
		}
		ans += tmp
	}
	return ans
}
