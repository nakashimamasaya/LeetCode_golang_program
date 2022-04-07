package _01046

import sort "sort"

// 1046. Last Stone Weight
// https://leetcode.com/problems/last-stone-weight/

func LastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Sort(sort.Reverse(sort.IntSlice(stones)))
		stones[1] = stones[0] - stones[1]
		stones = stones[1:]
	}

	return stones[0]
}
