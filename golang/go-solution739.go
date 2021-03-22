package main

func dailyTemperatures(T []int) []int {
	n := len(T)
	result := make([]int, n)
	for i := range T {
		count := 0
		for j := i + 1; j < n; j++ {
			if T[j] <= T[i] {
				count++
			} else {
				break
			}
		}
		result[i] = count
	}
	return result
}
