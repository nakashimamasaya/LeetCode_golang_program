package _00018

// 18. 4Sum
// https://leetcode.com/problems/4sum/

func fourSum(nums []int, target int) [][]int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	return searchFourSum(m, target, 0, keys, []int{})
}

func searchFourSum(numMap map[int]int, target, count int, keys, inprogress []int) [][]int {
	if target == 0 && count == 4 {
		return [][]int{inprogress}
	} else if count == 3 {
		for _, k := range keys {
			if k == target {
				return [][]int{append(inprogress, target)}
			}
		}
	}
	if count >= 3 {
		return [][]int{}
	}

	ans := [][]int{}

	for ki, k := range keys {
		tmp := []int{}
		for i := 1; i <= numMap[k] && count+i <= 4; i++ {
			tmp = append(tmp, k)
			ans = append(ans, searchFourSum(numMap, target-k*i, count+i, keys[ki+1:], append(inprogress, tmp...))...)
		}
	}
	return ans
}
