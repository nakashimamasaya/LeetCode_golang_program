package _0035

// 35. Search Insert Position
// https://leetcode.com/problems/search-insert-position/

func SearchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v >= target {
			return i
		}
	}
	return len(nums)
}
