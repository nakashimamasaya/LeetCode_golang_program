package _00234

import "reflect"

// 234. Palindrome Linked List
// https://leetcode.com/problems/palindrome-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	val := []int{}
	for head != nil {
		val = append(val, head.Val)
		head = head.Next
	}
	return reflect.DeepEqual(val[:len(val)/2], reverseSlice(val[len(val)-len(val)/2:]))
}

func reverseSlice(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}
