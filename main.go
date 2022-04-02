package main

import (
	"LeetCode/problem/00020"
	_0022 "LeetCode/problem/00022"
	"LeetCode/problem/00026"
	_0027 "LeetCode/problem/00027"
	_0028 "LeetCode/problem/00028"
	_0035 "LeetCode/problem/00035"
	_0058 "LeetCode/problem/00058"
	"LeetCode/problem/00704"
	"fmt"
)

func main() {
	// 20. Valid Parentheses
	separate("20")
	fmt.Println(_0020.IsValid("([)]"))

	// 22. Generate Parentheses
	separate("22")
	fmt.Println(_0022.GenerateParenthesis(4))

	// 26. Remove Duplicates from Sorted Array
	separate("26")
	fmt.Println(_0026.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	// 27. Remove Element
	separate("27")
	fmt.Println(_0027.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

	// 28. Implement strStr()
	separate("28")
	fmt.Println(_0028.StrStr("hello", "ll"))

	// 35. Search Insert Position
	separate("35")
	fmt.Println(_0035.SearchInsert([]int{1, 3, 5, 6}, 5))

	// 58. Length of Last Word
	separate("58")
	fmt.Println(_0058.LengthOfLastWord("   fly me   to   the moon  "))

	// 704. Binary Search
	separate("704")
	fmt.Println(_0704.Search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func separate(num string) {
	fmt.Println(num + "----------------------------------")
}
