package _0021

// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

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

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ans, ansH *ListNode

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		ans, ansH = list1, list1
		list1 = list1.Next
	} else {
		ans, ansH = list2, list2
		list2 = list2.Next
	}

	for list1N, list2N := list1, list2; list1N != nil || list2N != nil; {
		if list1N == nil {
			ans.Next = list2N
			list2N = nil
		} else if list2N == nil {
			ans.Next = list1N
			list1N = nil
		} else if list1N.Val < list2N.Val {
			ans.Next = list1N
			ans = list1N
			list1N = list1N.Next
		} else {
			ans.Next = list2N
			ans = list2N
			list2N = list2N.Next
		}
	}

	return ansH
}

func SetSample() *ListNode {
	one := ListNode{
		1,
		nil,
	}

	tmp := &one
	for i := 1; i < 2; i++ {
		num := ListNode{
			i,
			tmp,
		}
		tmp = &num
	}

	return tmp
}