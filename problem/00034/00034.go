package _00034

import "math"

// 34. Find First and Last Position of Element in Sorted Array
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}

	for i := 0; len(nums) > 0 && i <= len(nums)/2; i++ {
		if nums[i] == target && ans[0] == -1 {
			ans[0] = i
			ans[1] = int(math.Max(float64(i), float64(ans[1])))
		} else if nums[i] == target {
			ans[0] = int(math.Min(float64(i), float64(ans[0])))
			ans[1] = int(math.Max(float64(i), float64(ans[1])))
		}

		if nums[len(nums)-1-i] == target && ans[0] == -1 {
			ans[0] = len(nums) - 1 - i
			ans[1] = int(math.Max(float64(len(nums)-1-i), float64(ans[1])))
		} else if nums[len(nums)-1-i] == target {
			ans[0] = int(math.Min(float64(len(nums)-1-i), float64(ans[0])))
			ans[1] = int(math.Max(float64(len(nums)-1-i), float64(ans[1])))
		}
	}

	return ans
}
