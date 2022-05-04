package main

import (
	_0206 "LeetCode/problem/00206"
	"fmt"
)

func main() {
	// 206. Reverse Linked List
	separate("206")
	fmt.Println(_0206.ReverseList(_0206.SetSample()).Val)
}

func separate(num string) {
	fmt.Println(num + "----------------------------------")
}
