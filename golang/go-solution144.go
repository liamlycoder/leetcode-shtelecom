package main

func preorderTraversal(root *TreeNode) []int {
	var ret []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if node == nil {
			continue
		}
		ret = append(ret, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return ret
}
