package _00024

// 24. Swap Nodes in Pairs
// https://leetcode.com/problems/swap-nodes-in-pairs/

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

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	m := []*ListNode{}
	for ; head != nil; head = head.Next {
		m = append(m, head)
	}

	if len(m) == 1 {
		return m[0]
	}

	for i := 0; i < len(m)/2; i++ {
		m[i*2].Next = m[i*2+1].Next
		if i+1 < len(m)/2 {
			m[i*2].Next = m[i*2+2].Next
		}
		m[i*2+1].Next = m[i*2]
	}

	return m[1]
}
