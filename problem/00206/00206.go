package _0206

import "fmt"

// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/

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

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head.Next
	head.Next = nil

	for tmp != nil {
		tmp2 := tmp.Next
		tmp.Next = head
		head = tmp
		tmp = tmp2
	}

	return head
}

func SetSample() *ListNode {
	five := ListNode{
		5,
		nil,
	}

	fmt.Println(five)

	tmp := &five
	for i := 4; i > 0; i-- {
		num := ListNode{
			i,
			tmp,
		}
		tmp = &num
		fmt.Println(tmp)
	}

	return tmp
}
