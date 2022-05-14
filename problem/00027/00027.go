package _0027

import "sort"

// 27. Remove Element
// https://leetcode.com/problems/remove-element/

func RemoveElement(nums []int, val int) int {
	count := 0
	for i, v := range nums {
		if v == val {
			nums[i] = 100
		} else {
			count++
		}
	}

	sort.Ints(nums)
	return count
}
