package _02582

// 2582. Pass the Pillow
// https://leetcode.com/problems/pass-the-pillow/

func passThePillow(n int, time int) int {
	ans := []int{
		time%(n-1) + 1,
		n - time%(n-1),
	}
	return ans[time/(n-1)%2]
}
