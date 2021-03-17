package main

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for i := 0; len(queue) > 0; i++ {
		result = append(result, []int{})
		var temp []*TreeNode
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			result[i] = append(result[i], node.Val)
			if node.Left != nil{
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		queue = temp
	}
	return result
}

func main() {

}
