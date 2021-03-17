package main

func helper(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	left, right := helper(root.Left), helper(root.Right)

	selected := root.Val + left[1] + right[1]
	notSelected := max(left[0], left[1]) + max(right[1], right[0])

	return []int{selected, notSelected}
}


func rob3(root *TreeNode) int {
	result := helper(root)
	return max(result[0], result[1])
}


