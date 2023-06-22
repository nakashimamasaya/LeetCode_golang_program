package _00771

import "strings"

// 771. Jewels and Stones
// https://leetcode.com/problems/jewels-and-stones/

func numJewelsInStones(jewels string, stones string) int {
	ans := 0
	for i := 0; i < len(stones); i++ {
		if strings.Contains(jewels, string(stones[i])) {
			ans++
		}
	}

	return ans
}
