package _00283

// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	for i, start := 0, 1; i < len(nums)-start; i++ {
		for j := start; j < len(nums); j++ {
			if nums[j-1] == 0 {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				start = j
			}
		}
	}
}
