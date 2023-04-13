package _02399

// 2399. Check Distances Between Same Letters
// https://leetcode.com/problems/check-distances-between-same-letters/

func checkDistances(s string, distance []int) bool {
	const ASCII_a = 97
	d := make([]int, 26)

	for i, v := range s {
		if d[v-ASCII_a] == 0 {
			d[v-ASCII_a] = i + 1
		} else if i-d[v-ASCII_a] != distance[v-ASCII_a] {
			return false
		}
	}
	return true
}
