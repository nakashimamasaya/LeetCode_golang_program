package _0101

// 101. Symmetric Tree\n
// https://leetcode.com/problems/symmetric-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return checkTree(root.Left, root.Right)
}

func checkTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return checkTree(left.Left, right.Right) && checkTree(left.Right, right.Left)
}
