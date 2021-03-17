package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 先按照左端点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 保存结果
	var result [][]int

	// 保存试图合并的区间
	prev := intervals[0]

	// 从第二个区间开始，比较右端点和prev左端点的大小，尝试进行合并，不能合并就把prev直接加入到result中去，作为一个已经合并结束的区间了。
	for _, cur := range intervals {
		if prev[1] < cur[0] {   // 没有重叠
			result = append(result, prev)
			prev = cur
		} else {
			prev[1] = max(prev[1], cur[1])
		}
	}

	// 将最后一个区间也加入到结果中
	result = append(result, prev)
	return result
}

func main() {

}
