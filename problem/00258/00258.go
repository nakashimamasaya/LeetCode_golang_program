package _00258

// Add Digits
// https://leetcode.com/problems/add-digits/

func addDigits(num int) int {
	for num >= 10 {
		tmp := num
		for num = 0; tmp > 0; num, tmp = num+tmp%10, tmp/10 {
		}
	}
	return num
}
