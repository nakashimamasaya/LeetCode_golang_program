package _00876

// 876. Middle of the Linked List
// https://leetcode.com/problems/middle-of-the-linked-list/

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

func middleNode(head *ListNode) *ListNode {
	m := []*ListNode{}
	for ; head != nil; head = head.Next {
		m = append(m, head)
	}

	return m[len(m)/2]
}
