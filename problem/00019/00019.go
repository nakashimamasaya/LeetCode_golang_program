package _00019

// 19. Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	m := []*ListNode{}

	for ; head != nil; head = head.Next {
		m = append(m, head)
	}
	length, m := len(m), append(m, nil)
	if n == length {
		return m[1]
	}

	m[length-n-1].Next = m[length-n+1]
	return m[0]
}
