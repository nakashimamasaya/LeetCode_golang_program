package _00033

// 33. Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	for i := 0; i <= len(nums)/2; i++ {
		if nums[i] == target {
			return i
		} else if nums[len(nums)-1-i] == target {
			return len(nums) - 1 - i
		}
	}
	return -1
}
