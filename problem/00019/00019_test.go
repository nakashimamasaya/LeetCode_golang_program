package _00019

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	listVal := [][]int{
		{1, 2, 3, 4, 5},
		{1},
		{1, 2},
	}

	n := []int{
		2,
		1,
		1,
	}

	ansVal := [][]int{
		{1, 2, 3, 5},
		{},
		{1},
	}

	for i, a := range ansVal {
		head := ListNode{Val: listVal[i][0]}
		createNode(&head, listVal[i])
		if result := removeNthFromEnd(&head, n[i]); !reflect.DeepEqual(a, searchVal(result)) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i, searchVal(result), a)
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
	fmt.Println(head)
	val := []int{}
	for ; head != nil; head = head.Next {
		val = append(val, head.Val)
	}
	return val
}
