package _0021

import (
	"testing"
)

// 1 2 4   1 3 4   1 1 2 3 4 4
func TestSample1(t *testing.T) {
	ansVal := []int{1, 1, 2, 3, 4, 4}

	// create List1
	list1f := ListNode{Val: 4}
	list1t := ListNode{Val: 2, Next: &list1f}
	list1o := ListNode{Val: 1, Next: &list1t}

	// create List2
	list2f := ListNode{Val: 4}
	list2t := ListNode{Val: 3, Next: &list2f}
	list2o := ListNode{Val: 1, Next: &list2t}

	ans := MergeTwoLists(&list1o, &list2o)

	for i, list := 0, ans; list != nil; i, list = i+1, list.Next {
		if list.Val != ansVal[i] {
			t.Error("sample1 miss")
		}
	}

}

func TestSample2(t *testing.T) {
	ans := MergeTwoLists(nil, nil)

	if ans != nil {
		t.Error("sample2 miss")
	}
}

func TestSample3(t *testing.T) {
	list2 := ListNode{Val: 1}
	ans := MergeTwoLists(nil, &list2)

	if ans.Val != 1 || ans.Next != nil {
		t.Error("sample3 miss")
	}
}
