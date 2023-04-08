package _02578

import (
	"sort"
)

// 2578. Split With Minimum Sum
// https://leetcode.com/problems/split-with-minimum-sum/

func splitNum(num int) int {
	n := []int{}
	for ; num > 0; num /= 10 {
		n = append(n, num%10)
	}

	sort.Ints(n)

	for i := 2; i < len(n); i++ {
		n[i%2] = n[i%2]*10 + n[i]
	}

	return n[0] + n[1]
}
