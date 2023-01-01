package _02520

// 2520. Count the Digits That Divide a Number
// https://leetcode.com/problems/count-the-digits-that-divide-a-number/

func countDigits(num int) int {
	ans := 0
	for n := num; n > 0; n /= 10 {
		if num%(n%10) == 0 {
			ans++
		}
	}
	return ans
}
