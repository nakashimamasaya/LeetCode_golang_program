package _00015

import (
	"sort"
)

// 15. 3Sum
// https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1 && nums[i]+nums[j] <= 0; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k, n := j+1, (nums[i]+nums[j])*-1; k < len(nums) && nums[k] <= n; k++ {
				if n == nums[k] {
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
					break
				}
			}

		}
	}
	return ans
}
