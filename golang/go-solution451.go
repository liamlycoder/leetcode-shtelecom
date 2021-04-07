package main

import "sort"

func frequencySort(s string) string {
	// 采用数组来代替哈希表，效率略微高一点
	count := make([]int, 256)
	for _, v := range s {
		count[v]++
	}

	var bytes []byte
	for k, v := range count {
		if v > 0 {
			bytes = append(bytes, byte(k))
		}
	}

	sort.Slice(bytes, func(i, j int) bool {
		return count[bytes[i]] > count[bytes[j]]
	})

	var res []byte
	for i := 0; i < len(bytes); i++ {
		for j := 0; j < count[bytes[i]]; j++ {
			res = append(res, bytes[i])
		}
	}

	return string(res)
}

func main() {

}
