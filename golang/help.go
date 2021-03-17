package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}

func min(x, y int) int {
    if x > y {
        return y
    } else {
        return x
    }
}
