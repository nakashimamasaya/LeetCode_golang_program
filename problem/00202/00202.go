package _00202

import "math"

// 202. Happy Number
// https://leetcode.com/problems/happy-number/

func isHappy(n int) bool {
	return searchHappy(n, map[int]bool{})
}

func searchHappy(n int, hash map[int]bool) bool {
	if n == 1 {
		return true
	}

	if _, ok := hash[n]; ok {
		return false
	}

	cal := 0
	for tmp := n; tmp > 0; tmp /= 10 {
		cal += int(math.Pow(float64(tmp%10), 2))
	}
	
	hash[n] = true
	return searchHappy(cal, hash)
}
