package main


func pathSum(root *TreeNode, targetSum int) [][]int {
	var path []int
	var ans[][]int
	var dfs4 func(*TreeNode, int)   // 递归调用的时候，必须先声明该函数, 否则函数内部递归的时候找不到该函数
	dfs4 = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() {path = path[:len(path) - 1]}()
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs4(node.Left, left)
		dfs4(node.Right, left)
	}
	dfs4(root, targetSum)
	return ans
}
