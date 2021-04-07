package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n, best := len(nums), math.MaxInt32   // 这里我不知道为什么必须是int32才可以，int64就会报错

	// 根据差的绝对值来更新答案
	update := func(cur int) {
		if abs(cur - target) < abs(best - target) {
			best = cur
		}
	}

	// 枚举a
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		// 使用双指针来枚举b和c
		j, k := i + 1, n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				k0 := k - 1
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				j0 := j + 1
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1*x
	} else {
		return x
	}
}

func main() {


}
