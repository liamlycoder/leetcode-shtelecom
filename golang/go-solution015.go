package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	result := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		third := n - 1
		target := -nums[first]
		for second := first + 1; second < n; second++ {
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second] + nums[third] == target {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return result
}

func main() {
	nums := []int{54, 32, 9, 88, 534, 2}
	threeSum(nums)

}
