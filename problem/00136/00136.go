package _0136

import (
	"sort"
)

// 136. Single Number
// https://leetcode.com/problems/single-number/

func SingleNumber(nums []int) int {
	sort.Ints(nums)
	ans := nums[len(nums)-1]
	for i := 0; i+1 < len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			ans = nums[i]
			break
		}
	}
	return ans
}
