package _0104

import "math"

// 104. Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return int(searchDepth(1, root))
}

func searchDepth(depth float64, root *TreeNode) float64 {
	if root == nil {
		return depth - 1
	}

	return math.Max(searchDepth(depth+1, root.Left), searchDepth(depth+1, root.Right))
}
