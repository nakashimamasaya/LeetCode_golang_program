package main

import (
	"fmt"
)

func main() {

	// 704. Binary Search
	fmt.Println(search(
		[]int{-1, 0, 3, 5, 9, 12},
		9,
	))
}

// 704. Binary Search
// https://leetcode.com/problems/binary-search/
func search(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		} else if v > target {
			break
		}
	}
	return -1
}
