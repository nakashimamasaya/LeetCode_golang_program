package _0704

// 704. Binary Search
// https://leetcode.com/problems/binary-search/
func Search(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		} else if v > target {
			break
		}
	}
	return -1
}
