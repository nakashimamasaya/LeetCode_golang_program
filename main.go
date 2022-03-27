package main

import (
	"LeetCode/problem/00020"
	_0022 "LeetCode/problem/00022"
	"LeetCode/problem/00704"
	"fmt"
)

func main() {
	separate()

	// 20. Valid Parentheses
	fmt.Println(_0020.IsValid("([)]"))

	separate()

	// 22. Generate Parentheses
	fmt.Println(_0022.GenerateParenthesis(4))

	separate()
	// 704. Binary Search
	fmt.Println(_0704.Search([]int{-1, 0, 3, 5, 9, 12}, 9))
	separate()
}

func separate() {
	fmt.Println("----------------------------------")
}
