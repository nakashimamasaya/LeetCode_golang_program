package _0083

import "fmt"

// 83. Remove Duplicates from Sorted List
// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

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

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	prev := head
	tmp := head.Next

	for tmp != nil {
		if tmp.Val != prev.Val {
			prev.Next = tmp
			prev = tmp
		}
		tmp = tmp.Next
	}
	prev.Next = nil
	return head
}

func SetSample() *ListNode {
	one := ListNode{
		1,
		nil,
	}

	fmt.Println(one)

	tmp := &one
	for i := 1; i < 2; i++ {
		num := ListNode{
			i,
			tmp,
		}
		tmp = &num
		fmt.Println(tmp)
	}

	return tmp
}
