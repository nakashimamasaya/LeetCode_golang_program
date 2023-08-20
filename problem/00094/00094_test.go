package _00094

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	// -101 == null
	rootVal := [][]int{
		{1, -101, 2, 3},
		{},
		{1},
	}

	ans := [][]int{
		{1, 3, 2},
		{},
		{1},
	}

	for i, a := range ans {
		if result := inorderTraversal(createTree(rootVal[i])); !reflect.DeepEqual(a, result) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}

func createTree(rootVal []int) *TreeNode {
	if len(rootVal) == 0 {
		return nil
	}
	root := TreeNode{}
	root.Val = rootVal[0]
	tmpTree := &root
	for i := 1; i < len(rootVal); i++ {
		if i%2 == 1 && !nulljudgement(rootVal[i]) {
			tmpTree.Left = &TreeNode{Val: rootVal[i]}
		} else if i%2 == 0 && !nulljudgement(rootVal[i]) {
			tmpTree.Right = &TreeNode{Val: rootVal[i]}
		}

		if i%2 == 0 {
			if tmpTree.Left != nil {
				tmpTree = tmpTree.Left
			} else if tmpTree.Right != nil {
				tmpTree = tmpTree.Right
			}
		}
	}
	return &root
}

func nulljudgement(n int) bool {
	return n == -101
}
