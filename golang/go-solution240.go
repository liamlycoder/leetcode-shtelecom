package main

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])
	i, j := 0, col - 1
	for matrix[i][j] != target {
		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
		if i >= row || j < 0 {
			return false
		}
	}
	return true
}
