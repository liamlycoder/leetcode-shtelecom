package main

func flatten(root *TreeNode)  {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Left, cur.Right = nil, next
		}
		cur = cur.Right
	}

}

func main() {

}
