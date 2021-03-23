package main

// 最短无序连续子数组
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	stack := make([]int, n)
	l, r := n, 0
	for i, v := range nums {
		for len(stack) != 0 && nums[stack[len(stack) - 1]] > v {
			l = min(l, stack[len(stack) - 1])
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}
	stack = make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[stack[len(stack) - 1]] < nums[i] {
			r = max(r, stack[len(stack) - 1])
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}
	if r - l > 0 {
		return r - l + 1
	} else {
		return 0
	}
}
