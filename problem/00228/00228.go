package _00228

import "strconv"

// 228. Summary Ranges
// https://leetcode.com/problems/summary-ranges/

func summaryRanges(nums []int) []string {
	ans := []string{}
	for i, before := 0, 0; i < len(nums); i++ {
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			ans = append(ans, generateCaracter(nums[before], nums[i]))
			before = i + 1
		}
	}
	return ans
}

func generateCaracter(before, now int) string {
	if before == now {
		return strconv.Itoa(before)
	}
	return strconv.Itoa(before) + "->" + strconv.Itoa(now)
}
