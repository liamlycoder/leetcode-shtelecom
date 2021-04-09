package main

import "strconv"

/**
golang的slice切片不是纯引用类型
如果改变了原来的值，切片的值也会跟着变；如果改变的切片的值，原来的值也会变，这是引用类型的特性；
但是slice在函数传递的时候仍然是值传递，这一点不同于纯引用
 */
func binaryTreePaths(root *TreeNode) []string {
	var ans []string
	dfs3(root, "", &ans)
	return ans
}

func dfs3(root *TreeNode, path string, ans *[]string)  {
	if root == nil {
		return
	}
	path += strconv.Itoa(root.Val)
	if root.Left == root.Right {   // Leaf node
		*ans = append(*ans, path)
		return
	}
	dfs3(root.Left, path + "->", ans)
	dfs3(root.Right, path + "->", ans)
}
