package main

func uniquePaths(m int, n int) int {
	// dp[i][j]表示的是走到(i, j)的位置共有多少种不同的路径
	dp := make([][]int, m)

	// 第一行和第一列都初始化为1，因为只有一种走法
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// dp[i][j] = dp[i-1][j] + dp[1][j-1]  因为dp[i][j]的来源只有这两种路径的走法
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {

}
