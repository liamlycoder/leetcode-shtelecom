package main

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	f0, f1, f2 := -prices[0], 0, 0
	var newf0, newf1, newf2 int
	for i := 1; i < n; i++ {
		newf0 = max(f0, f2 - prices[i])
		newf1 = f0 + prices[i]
		newf2 = max(f1, f2)
		f0, f1, f2 = newf0, newf1, newf2
	}
	return max(f1, f2)
}


