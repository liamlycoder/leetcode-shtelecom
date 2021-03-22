package main

func countBits(num int) []int {
	result := make([]int, num + 1)
	for i := 1; i <= num; i++ {
		if i % 2 == 1 {
			result[i] = result[i - 1] + 1
		} else {
			result[i] = result[i / 2]
		}
	}
	return result
}
