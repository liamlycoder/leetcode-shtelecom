package main

func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mx*nums[i], min(nums[i], mn*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}


