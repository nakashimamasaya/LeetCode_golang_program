package _02404

// 2404. Most Frequent Even Element
// https://leetcode.com/problems/most-frequent-even-element/

func mostFrequentEven(nums []int) int {
	m := map[int]int{}
	ans := -1

	for _, n := range nums {
		m[n]++
	}

	for k, v := range m {
		if k%2 == 0 && (m[ans] < v || m[ans] == v && ans > k) {
			ans = k
		}
	}
	return ans
}
