package _01512

// 1512. Number of Good Pairs
// https://leetcode.com/problems/number-of-good-pairs/

func numIdenticalPairs(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}

	ans := 0
	for _, v := range m {
		ans += v * (v - 1) / 2
		// if v == 2 {
		// 	ans++
		// } else if v > 2 {

		// }
	}
	return ans
}
