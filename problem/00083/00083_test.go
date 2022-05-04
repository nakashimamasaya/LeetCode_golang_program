package _0083

import "testing"

func TestSample1(t *testing.T) {
	ansVal := []int{1, 2}

	listVal := []int{1, 1}
	list := ListNode{Val: 2, Next: nil}

	for _, val := range listVal {
		tmp := list
		list = ListNode{Val: val, Next: &tmp}
	}

	result := DeleteDuplicates(&list)

	for i, list := 0, result; list != nil; i, list = i+1, list.Next {
		if list.Val != ansVal[i] {
			t.Error("sample1 miss")
		}
	}
}

func TestSample2(t *testing.T) {
	ansVal := []int{1, 2, 3}

	listVal := []int{3, 2, 1, 1}
	list := ListNode{Val: 3, Next: nil}

	for _, val := range listVal {
		tmp := list
		list = ListNode{Val: val, Next: &tmp}
	}

	result := DeleteDuplicates(&list)

	for i, list := 0, result; list != nil; i, list = i+1, list.Next {
		if list.Val != ansVal[i] {
			t.Error("sample2 miss")
		}
	}
}
