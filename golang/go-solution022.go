package main

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(int, int, string)
	dfs = func(left int, right int, path string) {
		// 当构建长度达到2*n的时候，结束递归
		if 2*n == len(path) {
			res = append(res, path)
			return
		}

		// 只要有左括号，选择左括号都不算错，因为后续总是可能有右括号出来与之匹配的
		if left > 0 {
			dfs(left - 1, right, path + "(")
		}

		// 这里进行剪枝处理：当右括号比左括号少的时候，显然选用右括号是不合适的，直接减去
		if left < right {
			dfs(left, right - 1, path + ")")
		}
	}

	dfs(n, n, "")

	return res
}

func main() {

}
