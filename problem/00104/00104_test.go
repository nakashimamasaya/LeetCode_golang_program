package _0104

import "testing"

func TestSample1(t *testing.T) {
	root := TreeNode{Val: 3}
	createTree(&root, 9, 20)
	createTree(root.Right, 15, 7)

	ans := 3
	result := maxDepth(&root)

	if ans != result {
		t.Error("miss sample1")
	}
}

func TestSample2(t *testing.T) {
	root := TreeNode{Val: 1}
	createTree(&root, -200, 2)

	ans := 2
	result := maxDepth(&root)

	if ans != result {
		t.Error("miss sample2")
	}
}

func createTree(tree *TreeNode, leftVal int, rightVal int) {
	if leftVal != -200 {
		tree.Left = &TreeNode{Val: leftVal}
	}
	if rightVal != -200 {
		tree.Right = &TreeNode{Val: rightVal}
	}
}
