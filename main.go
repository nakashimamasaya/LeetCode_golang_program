package main

import (
	_0017 "LeetCode/problem/00017"
	_0206 "LeetCode/problem/00206"
	"fmt"
)

func main() {
	// 206. Reverse Linked List
	separate("206")
	fmt.Println(_0206.ReverseList(_0206.SetSample()).Val)

	separate("17")
	fmt.Println(_0017.LetterCombinations("2"))
}

func separate(num string) {
	fmt.Println(num + "----------------------------------")
}
