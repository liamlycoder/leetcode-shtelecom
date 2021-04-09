package main

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs2(root)
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func dfs2(node *TreeNode) (result int) {
	if node.Left != nil {
		if isLeafNode(node.Left) {
			result += node.Left.Val
		} else {
			result += dfs2(node.Left)
		}
	}
	if node.Right != nil && !isLeafNode(node.Right) {
		result += dfs2(node.Right)
	}
	return
}
