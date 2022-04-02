package _0066

// 66. Plus One
// https://leetcode.com/problems/plus-one/

func PlusOne(digits []int) []int {
	index := len(digits) - 1
	digits[index]++
	for index -= 1; index >= 0 && digits[index+1] == 10; index-- {
		digits[index]++
		digits[index+1] = 0
	}
	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}
	return digits
}
