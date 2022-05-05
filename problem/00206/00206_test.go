package _0206

import "testing"

func TestSample1(t *testing.T) {
	ansVal := []int{5, 4, 3, 2, 1}

	listVal := []int{4, 3, 2, 1}
	list := ListNode{Val: 5, Next: nil}

	for _, val := range listVal {
		tmp := list
		list = ListNode{Val: val, Next: &tmp}
	}

	result := ReverseList(&list)

	for i, list := 0, result; list != nil; i, list = i+1, list.Next {
		if list.Val != ansVal[i] {
			t.Error("sample1 miss")
		}
	}
}

func TestSample2(t *testing.T) {
	ansVal := []int{2, 1}

	listVal := []int{1}
	list := ListNode{Val: 2, Next: nil}

	for _, val := range listVal {
		tmp := list
		list = ListNode{Val: val, Next: &tmp}
	}

	result := ReverseList(&list)

	for i, list := 0, result; list != nil; i, list = i+1, list.Next {
		if list.Val != ansVal[i] {
			t.Error("sample2 miss")
		}
	}
}

func TestSample3(t *testing.T) {
	if ReverseList(nil) != nil {
		t.Errorf("miss sample3\nresult %v\nanswer %v", ReverseList(nil), nil)
	}
}
