package _0002

// 2. Add Two Numbers
// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func listNodeVal(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}

func listNodeNext(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := l1.Val + l2.Val
	list := ListNode{Val: sum % 10}
	l1, l2 = l1.Next, l2.Next
	for carry, tmpList := sum/10, &list; l1 != nil || l2 != nil || carry == 1; l1, l2, carry = listNodeNext(l1), listNodeNext(l2), sum/10 {
		l1Val, l2Val := listNodeVal(l1), listNodeVal(l2)
		sum = l1Val + l2Val + carry
		l := ListNode{Val: sum % 10}
		tmpList.Next = &l
		tmpList = &l
	}
	return &list
}
