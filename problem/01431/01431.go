package _01431

import "math"

// 1431. Kids With the Greatest Number of Candies
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/?envType=study-plan-v2&id=leetcode-75

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0.0
	for k, candie := range candies {
		maxCandies = math.Max(maxCandies, float64(candie))
		candies[k] += extraCandies
	}

	ans := []bool{}

	for _, candie := range candies {
		ans = append(ans, maxCandies <= float64(candie))
	}
	return ans
}
