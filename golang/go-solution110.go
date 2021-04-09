package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDepth := getDepth(root.Left)
	rightDepth := getDepth(root.Right)
	if abs(leftDepth - rightDepth) > 1 {
		return false
	} else {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getDepth(root.Left), getDepth(root.Right)) + 1
}

