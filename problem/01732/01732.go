package _01732

// 1732. Find the Highest Altitude
// https://leetcode.com/problems/find-the-highest-altitude

func largestAltitude(gain []int) int {
	ans, s := 0, 0

	for _, g := range gain {
		s += g
		if ans < s {
			ans = s
		}
	}

	return ans
}
