package _00075

// 75. Sort Colors
// https://leetcode.com/problems/sort-colors/

func sortColors(nums []int) {
	nCount := map[int]int{}

	for _, n := range nums {
		nCount[n]++
	}

	for i, index := 0, 0; i < 3; i++ {
		for j := 0; j < nCount[i]; j, index = j+1, index+1 {
			nums[index] = i
		}
	}
}
