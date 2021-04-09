package main

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	result := math.MaxInt32
	if root.Left != nil {
		result = min(minDepth(root.Left), result)
	}
	if root.Right != nil {
		result = min(minDepth(root.Right), result)
	}

	return result + 1
}

