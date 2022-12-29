package _02315

import "strings"

// 2315. Count Asterisks
// https://leetcode.com/problems/count-asterisks/

func countAsterisks(s string) int {
	array, ans := strings.Split(s, "|"), 0
	for i := 0; i < len(array); i += 2 {
		ans += strings.Count(array[i], "*")
	}
	return ans
}
