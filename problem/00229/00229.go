package _00229

// 229. Majority Element II
// https://leetcode.com/problems/majority-element-ii/description/

func majorityElement(nums []int) []int {
	var ans []int
	var m = map[int]int{}
	for _, n := range nums {
		m[n]++
		if m[n] > len(nums)/3 {
			ans, m[n] = append(ans, n), -50000
		}
	}
	return ans
}
