package main

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	var stack []*TreeNode

	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ret = append(ret, node.Val)
		node = node.Right
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}
	return ret
}
