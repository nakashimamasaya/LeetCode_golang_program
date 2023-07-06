package _00724

// 724. Find Pivot Index
// https://leetcode.com/problems/find-pivot-index

func pivotIndex(nums []int) int {
	sum, leftSum := 0, 0
	for _, n := range nums {
		sum += n
	}

	for i, n := range nums {
		sum -= n
		if sum == leftSum {
			return i
		}
		leftSum += n
	}
	return -1

}
