package main

func maxSubArray(nums []int) int {
	// 首先自己编写一个最大值函数
	var max func(nums...int)int
	max = func(nums ...int) int {
		res := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] > res {
				res = nums[i]
			}
		}
		return res
	}

	res := nums[0]
	pre := 0   // 到前一个索引为止，所得到的最大子序列之和
	for _, n := range nums {
		pre = max(pre + n, n)   // 实现状态转移方程
		res = max(res, pre)
	}
	return res
}

func main() {

}
