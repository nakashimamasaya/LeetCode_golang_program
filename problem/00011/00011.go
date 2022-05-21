package _0011

// 11. Container With Most Water
// https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
	ans := 0
	for i, h := range height {
		for j := len(height) - 1; j > i && ans < h*(j-i); j-- {
			ans = maxInt(ans, minInt(height[j], h)*(j-i))
			if height[j] >= h {
				break
			}
		}
	}
	return ans
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
