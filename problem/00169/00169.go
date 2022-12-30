package _00169

// 169. Majority Element
// https://leetcode.com/problems/majority-element/

func majorityElement(nums []int) int {
	m := map[int]int{}

	for _, n := range nums {
		m[n]++
		if m[n] > len(nums)/2 {
			return n
		}
	}
	return 0
}
