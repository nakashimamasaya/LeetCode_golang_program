package _0101

import "testing"

func TestSample1(t *testing.T) {
	root := TreeNode{Val: 1}
	createTree(&root, 2, 2)
	createTree(root.Left, 3, 4)
	createTree(root.Right, 4, 3)

	ans := true
	result := isSymmetric(&root)
	if ans != result {
		t.Error("miss sample1")
	}
}

func TestSample2(t *testing.T) {
	root := TreeNode{Val: 1}
	createTree(&root, 2, 2)
	createTree(root.Left, -200, 300)
	createTree(root.Left, -200, 300)

	ans := false
	result := isSymmetric(&root)
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
