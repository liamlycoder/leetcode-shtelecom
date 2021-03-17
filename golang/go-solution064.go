package main

func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])

	// dp[i][j]表示走到(i, j)位置的时候的路径最小和, 当然，这里直接把grid拿来当做dp数组，省去空间复杂度

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}

	return grid[m-1][n-1]
}

func main() {

}
