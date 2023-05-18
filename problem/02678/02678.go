package _02678

// 2678. Number of Senior Citizens
// https://leetcode.com/problems/number-of-senior-citizens/

func countSeniors(details []string) int {
	ans := 0
	for i := 0; i < len(details); i++ {
		if details[i][11] > 54 || details[i][11] == 54 && details[i][12] > 48 {
			ans++
		}
	}

	return ans
}
