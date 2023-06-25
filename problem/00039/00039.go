package _00039

// 39. Combination Sum
// https://leetcode.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	return searchSum(candidates, []int{}, 0, target)
}

func searchSum(candidates, status []int, now, target int) [][]int {
	if now == target {
		return [][]int{status}
	}

	ans := [][]int{}
	for i, v := range candidates {
		tmp := make([]int, len(status))
		copy(tmp, status)
		for j := 1; now+v*j <= target; j++ {
			tmp = append(tmp, v)
			ans = append(ans, searchSum(candidates[i+1:], tmp, now+v*j, target)...)
		}
	}
	return ans
}
