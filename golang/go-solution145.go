package main

func postorderTraversal(root *TreeNode) []int {
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
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	ret = reverse(ret)
	return ret
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

