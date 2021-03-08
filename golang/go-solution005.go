package main

func longestPalindrome(s string) string {
	n := len(s)
	result := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for length := 0; length < n; length++ {
		for i := 0; i + length < n; i++ {
			j := i + length
			if length == 0 {
				dp[i][j] = 1
			} else if length == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i + 1][j - 1]
				}
			}
			if dp[i][j] > 0 && length + 1 > len(result) {
				result = s[i: i+length+1]
			}
		}
	}
	return result
}

func main() {

}
