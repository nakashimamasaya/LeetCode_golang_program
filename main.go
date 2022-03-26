package main

import (
	"fmt"
)

func main() {
	separate()

	// 20. Valid Parentheses
	fmt.Println(isValid("([)]"))

	separate()

	// 704. Binary Search
	fmt.Println(search(
		[]int{-1, 0, 3, 5, 9, 12},
		9,
	))
	separate()
}

func separate() {
	fmt.Println("----------------------------------")
}

// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	var closeList []string
	closeMap := map[string]string{"{": "}", "(": ")", "[": "]"}
	for i := 0; i < len(s); i++ {
		if v, ok := closeMap[string(s[i])]; ok {
			closeList = append(closeList, v)
		} else if len(closeList) == 0 || string(s[i]) != closeList[len(closeList)-1] {
			return false
		} else {
			closeList = closeList[:len(closeList)-1]
		}
	}
	return len(closeList) == 0
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
