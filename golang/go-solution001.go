package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for index, num := range nums {
		if cp, ok := hash[target - num]; ok {
			return []int{index, cp}
		}
		hash[num] = index
	}
	return nil
}

func main() {
	nums := []int{3, 2, 3}
	res := twoSum(nums, 6)
	fmt.Println("res: ", res)
}
