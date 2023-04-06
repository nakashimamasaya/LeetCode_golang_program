package _01254

import (
	"math"
)

// 1254. Number of Closed Islands
// https://leetcode.com/problems/number-of-closed-islands/

func closedIsland(grid [][]int) int {
	ans := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 0 {
				ans += int(searchIsland(grid, i, j))
			}
		}
	}
	return ans
}

func searchIsland(grid [][]int, x, y int) float64 {
	if x <= 0 || y <= 0 || x >= len(grid)-1 || y >= len(grid[x])-1 {
		return 0
	}
	grid[x][y] = 2

	if grid[x-1][y] > 0 && grid[x+1][y] > 0 &&
		grid[x][y-1] > 0 && grid[x][y+1] > 0 {
		return 1
	}

	moveX := []int{1, -1, 0, 0}
	moveY := []int{0, 0, 1, -1}
	ans := 1.0
	for i := 0; i < 4; i++ {
		if grid[x+moveX[i]][y+moveY[i]] == 0 {
			ans = math.Min(ans, float64(searchIsland(grid, x+moveX[i], y+moveY[i])))
		}
	}

	return ans
}
