package _01342

// 1342. Number of Steps to Reduce a Number to Zero
// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/

func numberOfSteps(num int) int {
	ans := 0
	for ; num != 0; ans++ {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
	}
	return ans
}
