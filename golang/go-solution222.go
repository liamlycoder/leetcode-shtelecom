package main

// 普适性算法，非常简洁
//func countNodes(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	return countNodes(root.Left) + countNodes(root.Right) + 1
//}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := countCompleteTreeLevel(root.Left)
	right := countCompleteTreeLevel(root.Right)
	if left == right {
		return countNodes(root.Right) + (1<<left)
	} else {
		return countNodes(root.Left) + (1<<right)
	}
}

// 下面的算法只针对完全二叉树来统计层数
func countCompleteTreeLevel(root *TreeNode) int {
	level := 0
	for root != nil {
		level++
		root = root.Left
	}
	return level
}

