package main


func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var dfs func(start int, temp[]int, sum int)   // 以start索引为起点，temp为已加入的数，sum为当前的数字和，进行搜索
	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				newTemp := make([]int, len(temp))
				copy(newTemp, temp)
				res = append(res, newTemp)
			}
			return
		}

		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])   // 尝试把当前数字加进入
			dfs(i, temp, sum + candidates[i])    // 尝试进行搜索
			temp = temp[: len(temp) - 1]   // 回溯
		}
	}

	dfs(0, []int{}, 0)
	return res
}

func main() {

}
