package _02300

import (
	"math"
)

// 2300. Successful Pairs of Spells and Potions
// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/

func successfulPairs(spells []int, potions []int, success int64) []int {
	cal := make([]int, 100000)
	for _, potion := range potions {
		c := math.Ceil(float64(success)/float64(potion)) - 1
		if c < 100000 {
			cal[int(c)]++
		}
	}

	for i, c := 0, 0; i < 100000; i++ {
		c += cal[i]
		cal[i] = c
	}

	ans := make([]int, len(spells))
	for i, spell := range spells {
		ans[i] = cal[spell-1]
	}

	return ans
}
