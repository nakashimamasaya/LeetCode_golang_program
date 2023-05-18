package _00024

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	listVal := [][]int{
		{1, 2, 3, 4},
		{},
		{1},
	}

	ansVal := [][]int{
		{2, 1, 4, 3},
		{},
		{1},
	}

	for i, a := range ansVal {
		var tmp *ListNode
		if len(listVal[i]) == 0 {
			tmp = nil
		} else {
			tmp = &ListNode{Val: listVal[i][0]}
			createNode(tmp, listVal[i])
		}
		if result := swapPairs(tmp); !reflect.DeepEqual(a, searchVal(result)) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, searchVal(result), a)
		}

	}
}

func createNode(head *ListNode, val []int) {
	for i := 1; i < len(val); i++ {
		tmp := ListNode{Val: val[i]}
		head.Next = &tmp
		head = head.Next
	}
}

func searchVal(head *ListNode) []int {
	val := []int{}
	for ; head != nil; head = head.Next {
		val = append(val, head.Val)
	}
	return val
}
