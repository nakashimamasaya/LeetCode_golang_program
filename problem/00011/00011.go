package _0011

// 11. Container With Most Water
// https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
	ans := 0

	for i, j, tmp := 0, len(height)-1, len(height)-1; i < j; tmp = j - i {
		if height[i] > height[j] {
			tmp, j = height[j]*tmp, j-1
		} else {
			tmp, i = height[i]*tmp, i+1
		}
		if ans < tmp {
			ans = tmp
		}
	}
	return ans
}
