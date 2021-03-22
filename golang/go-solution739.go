package main

func dailyTemperatures(T []int) []int {
	n := len(T)
	result := make([]int, n)
	var stack []int

	for i := 0; i < n; i++ {
		temp := T[i]
		for len(stack) > 0 && temp > T[stack[len(stack) - 1]] {
			prevIndex := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			result[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return result
}
