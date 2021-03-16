package main

// 思路：可以将二叉树的直径转换为：二叉树的 每个节点的左右子树的高度和的最大值。

// 该函数其实有两个功能：1）返回最大深度；2）更新maxDepth，使其保存当前左右子树的最大高度和
func depth(root *TreeNode, maxDepth *int) int {
	if root == nil {
		return 0
	}
	a := depth(root.Left, maxDepth)
	b := depth(root.Right, maxDepth)
	*maxDepth = max(a + b, *maxDepth)
	return max(a, b) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	m := 0
	maxDepth := &m
	depth(root, maxDepth)
	return *maxDepth
}

func main() {

}
