package main

func rotate(matrix [][]int)  {
	n := len(matrix)

	// 上下水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n - i - 1] = matrix[n - i - 1], matrix[i]
	}

	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main() {

}
