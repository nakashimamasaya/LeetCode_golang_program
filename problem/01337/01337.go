package _01337

import "sort"

// 1337. The K Weakest Rows in a Matrix
// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

func kWeakestRows(mat [][]int, k int) []int {
	ans, weakest, sol := []int{}, map[int][]int{}, []int{}

	for i, m := 0, 0; i < len(mat); i, m = i+1, 0 {
		for _, mm := range mat[i] {
			m += mm
		}
		weakest[m], sol = append(weakest[m], i), append(sol, m)
	}

	sort.Ints(sol)

	for i := 0; i < k; i += len(weakest[sol[i]]) {
		ans = append(ans, weakest[sol[i]]...)
	}
	return ans[0:k]
}
