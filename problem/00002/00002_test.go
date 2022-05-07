package _0002

import (
	"testing"
)

func TestSample1(t *testing.T) {
	l1Val, l2Val := []int{4, 2}, []int{6, 5}
	l1, l2 := createList(l1Val, &ListNode{Val: 3}), createList(l2Val, &ListNode{Val: 4})

	result := addTwoNumbers(l1, l2)
	ansVal := []int{7, 0, 8}

	for i := 0; result != nil; result, i = result.Next, i+1 {
		if result.Val != ansVal[i] {
			t.Error("miss sample1")
		}
	}
}

func TestSample2(t *testing.T) {
	l1, l2 := ListNode{Val: 0}, ListNode{Val: 0}

	result := addTwoNumbers(&l1, &l2)
	ansVal := []int{0}

	for i := 0; result != nil; result, i = result.Next, i+1 {
		if result.Val != ansVal[i] {
			t.Error("miss sample2")
		}
	}
}

func TestSample3(t *testing.T) {
	l1Val, l2Val := []int{9, 9, 9, 9, 9, 9}, []int{9, 9, 9}
	l1, l2 := createList(l1Val, &ListNode{Val: 9}), createList(l2Val, &ListNode{Val: 9})

	result := addTwoNumbers(l1, l2)
	ansVal := []int{8, 9, 9, 9, 0, 0, 0, 1}

	for i := 0; result != nil; result, i = result.Next, i+1 {
		if result.Val != ansVal[i] {
			t.Error("miss sample3")
		}
	}
}

func createList(val []int, l *ListNode) *ListNode {
	for _, v := range val {
		tmp := l
		l = &ListNode{Val: v, Next: tmp}
	}
	return l
}
