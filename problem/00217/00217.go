package _0217

import "sort"

// 217. Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/

func ContainsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
