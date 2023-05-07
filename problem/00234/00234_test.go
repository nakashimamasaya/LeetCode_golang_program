package _00234

import "testing"

func TestSample(t *testing.T) {
	head := [][]int{
		{1, 2, 2, 1},
		{1, 2},
	}

	ans := []bool{
		true,
		false,
	}

	for i, a := range ans {
		if result := isPalindrome(createList(head[i])); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}

}

func createList(val []int) *ListNode {
	firstHead := ListNode{Val: val[0]}
	head := &firstHead
	for i := 1; i < len(val); i++ {
		tmp := ListNode{Val: val[i]}
		head.Next = &tmp
		head = &tmp
	}

	return &firstHead
}
