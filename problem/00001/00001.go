package _0001

import "sort"

// 1. Two Sum
// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	tmp := append([]int{}, nums...)
	sort.Ints(nums)
	for i, v := range nums {
		for j := len(nums) - 1; j > i; j-- {
			if v+nums[j] == target {
				first := findIndex(tmp, v, -1)
				second := findIndex(tmp, nums[j], first)
				if first == -1 || second == -1 {
					continue
				}
				return []int{first, second}
			}
		}
	}
	return []int{}
}

func findIndex(nums []int, target int, skip int) int {
	for i, v := range nums {
		if v == target && i != skip {
			return i
		}
	}
	return -1
}
