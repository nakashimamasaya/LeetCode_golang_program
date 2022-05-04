package main

import (
	_0083 "LeetCode/problem/00083"
	_0206 "LeetCode/problem/00206"
	_0704 "LeetCode/problem/00704"
	"fmt"
)

func main() {
	// 83. Remove Duplicates from Sorted List
	separate("83")
	fmt.Println(_0083.DeleteDuplicates(_0083.SetSample()))

	// 206. Reverse Linked List
	separate("206")
	fmt.Println(_0206.ReverseList(_0206.SetSample()).Val)

	// 704. Binary Search
	separate("704")
	fmt.Println(_0704.Search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func separate(num string) {
	fmt.Println(num + "----------------------------------")
}
