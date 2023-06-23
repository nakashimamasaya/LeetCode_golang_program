package _00334

import (
	"sort"
)

// 334. Increasing Triplet Subsequence
// https://leetcode.com/problems/increasing-triplet-subsequence

func increasingTriplet(nums []int) bool {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	if len(m) < 3 {
		return false
	}

	for i, iV := range nums[:len(nums)-2] {
		for j, jV := range nums[i+1 : len(nums)-1] {
			if iV >= jV {
				continue
			}
			tmpS := make([]int, len(nums[i+j+1:]))
			n := copy(tmpS, nums[i+j+1:])
			sort.Ints(tmpS)
			if jV < tmpS[n-1] {
				return true
			}
		}
	}
	return false
}
