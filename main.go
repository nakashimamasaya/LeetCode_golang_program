package main

import (
	_0083 "LeetCode/problem/00083"
	_0171 "LeetCode/problem/00171"
	_0206 "LeetCode/problem/00206"
	_0704 "LeetCode/problem/00704"
	_01046 "LeetCode/problem/01046"
	"fmt"
)

func main() {
	// 83. Remove Duplicates from Sorted List
	separate("83")
	fmt.Println(_0083.DeleteDuplicates(_0083.SetSample()))

	// 171. Excel Sheet Column Number
	separate("171")
	fmt.Println(_0171.TitleToNumber("AB"))

	// 206. Reverse Linked List
	separate("206")
	fmt.Println(_0206.ReverseList(_0206.SetSample()).Val)

	// 704. Binary Search
	separate("704")
	fmt.Println(_0704.Search([]int{-1, 0, 3, 5, 9, 12}, 9))

	// 1046. Last Stone Weight
	separate("1046")
	fmt.Println(_01046.LastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}

func separate(num string) {
	fmt.Println(num + "----------------------------------")
}
