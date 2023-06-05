package _02652

// 2652. Sum Multiples
// https://leetcode.com/problems/sum-multiples/

func sumOfMultiples(n int) int {
	ans := 0
	for ; n >= 3; n-- {
		if n%3 == 0 || n%5 == 0 || n%7 == 0 {
			ans += n
		}
	}
	return ans
}
