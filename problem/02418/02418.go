package _02418

import "sort"

// 2418. Sort the People
// https://leetcode.com/problems/sort-the-people/

func sortPeople(names []string, heights []int) []string {
	heightMap := map[int]string{}
	for i := 0; i < len(names); i++ {
		heightMap[heights[i]] = names[i]
	}

	var ans []string
	sort.Sort(sort.Reverse(sort.IntSlice(heights)))
	for _, h := range heights {
		ans = append(ans, heightMap[h])
	}
	return ans
}
