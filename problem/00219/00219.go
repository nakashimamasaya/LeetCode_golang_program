package _0219

// 219. Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii/

func ContainsNearbyDuplicate(nums []int, k int) bool {
	for i, v := range nums {
		for j := i + 1; j-i <= k && j < len(nums); j++ {
			if v == nums[j] {
				return true
			}
		}
	}
	return false
}
