package _0026

import (
	"sort"
)

// 26. Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func RemoveDuplicates(nums []int) int {
	var resultMap = map[int]bool{}
	for i, num := range nums {
		if _, ok := resultMap[num]; ok {
			nums[i] = 300
		}
		resultMap[num] = true
	}

	sort.Ints(nums)
	nums = nums[:len(resultMap)]

	return len(nums)
}
